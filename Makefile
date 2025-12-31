# ==========================
# Configurações gerais
# ==========================
EBOOK_DIR := ebook
BOOK_NAME := workshop
SOURCE_FILE := $(EBOOK_DIR)/workshop.md
OUTPUT_PDF := $(EBOOK_DIR)/$(BOOK_NAME).pdf

# ==========================
# Docker / Pandoc
# ==========================
PDF_BUILDER = docker
PDF_BUILDER_FLAGS = \
	run \
	--rm \
	-u `id -u`:`id -g` \
	-e PYTHONWARNINGS=ignore::SyntaxWarning \
	-v `pwd`:/pandoc \
	-w /pandoc/$(EBOOK_DIR) \
	dalibo/pandocker \
	--pdf-engine=xelatex \
	--template=eisvogel.tex \
	--filter pandoc-include \
	--filter pandoc-latex-environment \
	--listings \
	--top-level-division=chapter

# ==========================
# Build
# ==========================
.PHONY: build
build:
	@echo "Construindo o e-book em PDF..."

	@sed -Ei "s/[0-3][0-9]-[01][0-9]-[12][0-9]{3}/$(shell date +%d-%m-%Y)/" $(SOURCE_FILE)

	@cp dia_01/README.md $(EBOOK_DIR)/dia_01.md
	@cp dia_02/README.md $(EBOOK_DIR)/dia_02.md
	@cp dia_03/README.md $(EBOOK_DIR)/dia_03.md
	@cp -r dia_03/img $(EBOOK_DIR)/img

	@$(PDF_BUILDER) $(PDF_BUILDER_FLAGS) workshop.md -o $(BOOK_NAME).pdf

	@rm -f \
		$(EBOOK_DIR)/dia_01.md \
		$(EBOOK_DIR)/dia_02.md \
		$(EBOOK_DIR)/dia_03.md
	@rm -rf $(EBOOK_DIR)/img

	@echo "Concluído: $(OUTPUT_PDF)"

# ==========================
# Clean
# ==========================
.PHONY: clean
clean:
	rm -f $(EBOOK_DIR)/$(BOOK_NAME).pdf
	rm -f $(EBOOK_DIR)/$(BOOK_NAME).epub
	rm -f $(EBOOK_DIR)/$(BOOK_NAME).mobi

.PHONY: update-readme
update-readme:
	@go run tools/update_readme.go
