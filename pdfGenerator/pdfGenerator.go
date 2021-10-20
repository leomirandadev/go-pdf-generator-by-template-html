package pdfGenerator

type PDFGenetorInterface interface {
	Create(htmlFile string) (string, error)
}
