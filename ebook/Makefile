SOURCE_FILE_NAME = workshop.md
BOOK_FILE_NAME = workshop

PDF_BUILDER = docker
PDF_BUILDER_FLAGS = \
	run \
	--rm \
	-u `id -u`:`id -g` \
	-v `pwd`:/pandoc \
	dalibo/pandocker \
	--pdf-engine=xelatex \
	--template=eisvogel.tex \
	--filter pandoc-include \
	--listings \
	--top-level-division=chapter

build:
	@echo "Construindo o e-book em PDF..."
	@sed -Ei "s/[0-3][0-9]-[01][0-9]-[12][0-9]{3}/$(shell date +%d-%m-%Y%n)/" $(SOURCE_FILE_NAME);
	@cp ../dia_01/README.md ./dia_01.md && \
	cp ../dia_02/README.md ./dia_02.md && \
	cp ../dia_03/README.md ./dia_03.md && \
	cp -r ../dia_03/img . && $(PDF_BUILDER) $(PDF_BUILDER_FLAGS) $(SOURCE_FILE_NAME) -o $(BOOK_FILE_NAME).pdf
	@rm -r dia_01.md dia_02.md dia_03.md img
	@echo "Concluído"

clean:
	rm -f $(BOOK_FILE_NAME).pdf
	rm -f $(BOOK_FILE_NAME).epub
	rm -f $(BOOK_FILE_NAME).mobi
