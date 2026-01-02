local NL = pandoc.RawInline("tex", "\\newline ")

local function normalize_nbsp(s)
  if s == nil then return "" end
  s = s:gsub("&nbsp;", "\194\160")
  s = s:gsub("&#160;", "\194\160")
  return (s:gsub("\194\160", " "))
end

local function normalize_smart_quotes(s)
  if s == nil then return "" end
  s = s:gsub("“", '"'):gsub("”", '"')
  s = s:gsub("‘", "'"):gsub("’", "'")
  return s
end

local function pick_delim(s)
  local delims = {"|", "!", "@", "#", "%", ";", ":", "~", "^", "+"}
  for _, d in ipairs(delims) do
    if not s:find(d, 1, true) then return d end
  end
  return "|"
end

-- Extrai linguagem de <code class="language-go">, <code class="lang-go"> etc.
local function parse_code_language(tag)
  local cls = tag:match('class%s*=%s*"([^"]+)"')
  if not cls then
    return "__PLAIN__" -- fallback quando não há class
  end

  local lang = cls:match("language%-([%w_%-]+)") or cls:match("lang%-([%w_%-]+)")
  if not lang then
    return "__PLAIN__" -- fallback quando class existe mas não tem language-*
  end

  lang = lang:lower()

  -- "text" => plain text
  if lang == "text" or lang == "plain" or lang == "plaintext" then
    return "__PLAIN__"
  end

  -- Mapeamento para nomes do listings (ajuste se necessário)
  if lang == "go" or lang == "golang" then return "Go" end
  if lang == "json" then return "json" end
  if lang == "yaml" or lang == "yml" then return "yaml" end
  if lang == "bash" or lang == "sh" then return "bash" end
  if lang == "sql" then return "SQL" end

  -- fallback: tenta passar como está (listings pode ou não reconhecer)
  return lang
end

local function to_lstinline(text, opts)
  text = normalize_nbsp(text)
  text = normalize_smart_quotes(text)

  local d = pick_delim(text)
  if d == "|" then text = text:gsub("|", "\\|") end

  -- Base
  local optstr = "keepspaces=true"

  -- Regra:
  -- - Se opts.language existir e NÃO for __PLAIN__, aplicamos language=<...>
  -- - Caso contrário (nil / __PLAIN__), forçamos style=plaincode e zeramos language
  if opts and opts.language and opts.language ~= "__PLAIN__" then
    optstr = optstr .. ",language=" .. opts.language
  else
    optstr = optstr .. ",style=plaincode"
  end

  return pandoc.RawInline("tex", "\\lstinline[" .. optstr .. "]" .. d .. text .. d)
end

-- Processa inlines dentro de célula: trata <code>...</code> multilinha preservando <br>/LineBreak
local function process_inlines(inlines)
  local out = pandoc.List()

  local in_code = false
  local current_line = {}
  local code_opts = nil

  local function flush_line()
    local s = table.concat(current_line)
    out:insert(to_lstinline(s, code_opts))
    current_line = {}
  end

  local function append_text(s)
    table.insert(current_line, s)
  end

  for _, el in ipairs(inlines) do
    if el.t == "RawInline" and el.format == "html" then
      -- abertura <code ...>
      if el.text:match("^<code[^>]*>$") then
        in_code = true
        current_line = {}
        code_opts = { language = parse_code_language(el.text) }

      -- fechamento </code>
      elseif el.text:match("^</code>$") then
        flush_line()
        in_code = false
        code_opts = nil

      -- <br> dentro do code
      elseif in_code and el.text:match("^<br%s*/?>$") then
        flush_line()
        out:insert(NL)

      -- <br> fora do code
      elseif (not in_code) and el.text:match("^<br%s*/?>$") then
        out:insert(NL)

      else
        if not in_code then
          out:insert(el)
        end
      end

    elseif in_code then
      -- Dentro do <code>...</code>
      if el.t == "LineBreak" then
        flush_line()
        out:insert(NL)
      elseif el.t == "SoftBreak" then
        append_text(" ")
      elseif el.t == "Space" then
        append_text(" ")
      else
        append_text(pandoc.utils.stringify(el))
      end

    else
      -- Fora do <code>...</code>
      if el.t == "LineBreak" then
        out:insert(NL)
      elseif el.t == "Code" then
        -- Sem linguagem explícita -> plain preto
        out:insert(to_lstinline(el.text, { language = "__PLAIN__" }))
      else
        out:insert(el)
      end
    end
  end

  if in_code then
    flush_line()
  end

  return out
end

local function fix_cell(cell)
  for i, blk in ipairs(cell.contents) do
    cell.contents[i] = pandoc.walk_block(blk, {
      Inlines = function(inl)
        return process_inlines(inl)
      end
    })
  end
end

function Table(el)
  if el.head and el.head.rows then
    for _, row in ipairs(el.head.rows) do
      for _, cell in ipairs(row.cells) do
        fix_cell(cell)
      end
    end
  end

  if el.bodies then
    for _, body in ipairs(el.bodies) do
      for _, row in ipairs(body.body) do
        for _, cell in ipairs(row.cells) do
          fix_cell(cell)
        end
      end
    end
  end

  if el.foot and el.foot.rows then
    for _, row in ipairs(el.foot.rows) do
      for _, cell in ipairs(row.cells) do
        fix_cell(cell)
      end
    end
  end

  return el
end
