#!/usr/bin/env python3
# -*- coding: utf-8 -*-

"""
Gera/atualiza automaticamente o README.md principal do projeto
a partir dos README.md dos diretórios dia_XX.

Regras adotadas:
- # Dia XX        -> ignorado (já representado no README principal)
- ## Título       -> item numerado
- ### Subtítulo   -> subitem indentado
- Âncoras seguem o padrão aproximado do GitHub (PT-BR friendly)
"""

import re
import unicodedata
from pathlib import Path


ROOT = Path(".")
OUT = ROOT / "README.md"


def gh_slug(text: str) -> str:
    """
    Gera âncora compatível com o GitHub:
    - lowercase
    - mantém acentos
    - remove pontuação
    - espaços viram hífens
    """
    text = text.strip().lower()

    # remove pontuação, mas mantém letras acentuadas
    text = re.sub(r"[^\w\s\-áéíóúàèìòùâêîôûãõç]", "", text, flags=re.UNICODE)

    # espaços -> hífen
    text = re.sub(r"\s+", "-", text)

    # colapsa múltiplos hífens
    text = re.sub(r"-{2,}", "-", text).strip("-")

    return text


def read_headings(md_path: Path):
    """
    Lê headings Markdown (##, ###, etc.), ignorando blocos de código.
    Retorna lista de tuplas: (nível, título).
    """
    headings = []
    in_code_block = False

    for line in md_path.read_text(encoding="utf-8").splitlines():
        if line.strip().startswith("```"):
            in_code_block = not in_code_block
            continue

        if in_code_block:
            continue

        match = re.match(r"^(#{1,6})\s+(.+?)\s*$", line)
        if not match:
            continue

        level = len(match.group(1))
        title = match.group(2).strip()
        headings.append((level, title))

    return headings


def day_label_and_anchor(day_dir_name: str):
    """
    Converte:
      dia_02 -> ("DIA 02", "dia-02")
    """
    number = day_dir_name.split("_")[1]
    label = f"DIA {number}"
    anchor = gh_slug(label)
    return label, anchor


def main():
    day_dirs = sorted(
        [
            p for p in ROOT.iterdir()
            if p.is_dir() and re.match(r"^dia_\d{2}$", p.name)
        ],
        key=lambda p: p.name,
    )

    lines = []
    lines.append("# Workshop Go: do Zero à API")
    lines.append("")

    global_index = 1

    for day_dir in day_dirs:
        day_label, day_anchor = day_label_and_anchor(day_dir.name)
        day_link = f"../../tree/main/{day_dir.name}#{day_anchor}"

        lines.append(f"## [{day_label}]({day_link})")
        lines.append("")

        day_readme = day_dir / "README.md"
        if not day_readme.exists():
            lines.append("_README.md não encontrado para este dia._")
            lines.append("")
            continue

        headings = read_headings(day_readme)

        has_current_item = False
        sub_index = 0

        for level, title in headings:
            # Ignora "# Dia XX"
            if level == 1:
                continue

            # Item principal
            if level == 2:
                anchor = gh_slug(title)
                link = f"../../tree/main/{day_dir.name}#{anchor}"
                lines.append(f"{global_index}. [{title}]({link})")
                lines.append("")
                global_index += 1
                sub_index = 0
                has_current_item = True
                continue

            # Subitem
            if level == 3 and has_current_item:
                sub_index += 1
                anchor = gh_slug(title)
                link = f"../../tree/main/{day_dir.name}#{anchor}"
                lines.append(f"    {sub_index}. [{title}]({link})")
                continue

        lines.append("")

    OUT.write_text("\n".join(lines).rstrip() + "\n", encoding="utf-8")
    print("README.md atualizado com sucesso.")


if __name__ == "__main__":
    main()
