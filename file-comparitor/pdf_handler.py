from PyPDF2 import PdfReader

class PDFHandler:
    @staticmethod
    def read_pdf(file_path):
        """
        Read and extract text from a PDF file
        """
        try:
            reader = PdfReader(file_path)
            text = ""
            for page in reader.pages:
                text += page.extract_text() + "\n"
            return text, None
        except Exception as e:
            return None, str(e)
