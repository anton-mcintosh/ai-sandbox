import tkinter as tk
from tkinter import ttk, filedialog, messagebox
from langchain_ollama import OllamaLLM
from pdf_handler import PDFHandler
import os

class ChatApp:
    def __init__(self, root, model="llama2"):
        self.root = root
        self.root.title("AI Chat Interface")
        
        # Initialize Ollama and PDF handler
        self.llm = OllamaLLM(model=model)
        self.pdf_handler = PDFHandler()
        
        # Initialize conversation history
        self.conversation_history = []
        
        # Initialize PDF paths
        self.req_path = None
        self.prop_path = None
        
        self.setup_ui()

    def setup_ui(self):
# Create main container
        main_frame = ttk.Frame(self.root, padding="10")
        main_frame.pack(fill=tk.BOTH, expand=True)
        
        # Create and configure text display area
        self.output_area = tk.Text(main_frame, height=50, width=50, wrap=tk.WORD)
        self.output_area.pack(fill=tk.BOTH, expand=True)
        
        # Add scrollbar
        scrollbar = ttk.Scrollbar(main_frame, orient='vertical', command=self.output_area.yview)
        scrollbar.pack(side=tk.RIGHT, fill=tk.Y)
        self.output_area['yscrollcommand'] = scrollbar.set
        
        # Make the text read-only
        self.output_area.config(state='disabled')
        
        # Create input field
        self.input_field = ttk.Entry(main_frame, width=40)
        self.input_field.pack(fill=tk.X, pady=5)

        # Create file label
        self.file_label = ttk.Label(main_frame, text="No file selected")
        self.file_label.pack(pady=5self.root = root)

        # button for requirements pdf
        self.attach_pdf_button = ttk.Button(main_frame, text="Attach Requirements PDF", command=self.attach_pdf)
        self.attach_pdf_button.pack(pady=5)
        
        # button for proposal pdf
        self.attach_proposal_button = ttk.Button(main_frame, text="Attach Proposal PDF", command=self.attach_proposal)
        self.attach_proposal_button.pack(pady=5)
        
        # Create send button
        send_button = ttk.Button(main_frame, text="Send", command=self.send_message)
        send_button.pack(pady=5)
        
        # Bind Enter key to send message
        self.input_field.bind("<Return>", lambda e: self.send_message())

    def select_pdf(self, pdf_num):
        file_path = filedialog.askopenfilename(filetypes=[("PDF files", "*.pdf")])
        if file_path:
            if pdf_num == 1:
                self.pdf1_path = file_path
                self.pdf1_label.config(text=os.path.basename(file_path))
            else:
                self.pdf2_path = file_path
                self.pdf2_label.config(text=os.path.basename(file_path))

    def compare_pdfs(self):
        if not self.pdf1_path or not self.pdf2_path:
            messagebox.showerror("Error", "Please select both PDFs first")
            return

        # Get the context
        context = self.context_input.get("1.0", tk.END).strip()
        
        # Process PDFs and context
        pdf1_text = self.process_pdf(self.pdf1_path)
        pdf2_text = self.process_pdf(self.pdf2_path)
        
        # Construct the message
        message = f"Compare these two PDFs:\n\nPDF 1:\n{pdf1_text}\n\nPDF 2:\n{pdf2_text}"
        if context:
            message += f"\n\nAdditional Context:\n{context}"
        
        # Add message to chat history
        self.add_to_chat("User", message)
        
        # Get AI response
        response = self.get_ai_response(message)
        self.add_to_chat("Assistant", response)

    def process_pdf(self, pdf_path):
        # Read PDF content using the handler
        pdf_text, error = self.pdf_handler.read_pdf(pdf_path)
        
        if error:
            self.file_label.config(text=f"Error reading PDF: {error}")
            return ""
        
        # Update file label
        filename = pdf_path.split('/')[-1]
        self.file_label.config(text=f"Selected: {filename}")
        
        # Enable text widget for updating
        self.output_area.config(state='normal')
        
        # Add PDF content to conversation
        self.conversation_history.append(f"Human: I'm sharing a PDF with the following content:\n{pdf_text}")
        self.output_area.insert(tk.END, f"[PDF uploaded: {filename}]\n\n")
        
        # Get AI response about the PDF
        # full_context = "\n".join(self.conversation_history)
        # response = self.llm.invoke(full_context)
        # cleaned_response = self.clean_response(response)
        
        # Add AI response to history
        self.conversation_history.append(f"Assistant: {pdf_text}")
        
        # Display AI response
        self.output_area.insert(tk.END, f"AI: {pdf_text}\n\n")
        
        # Make text widget read-only again
        self.output_area.config(state='disabled')
        
        # Scroll to bottom
        self.output_area.see(tk.END)
        
        return pdf_text

    def send_message(self):
        user_input = self.input_field.get()
        if user_input.strip():
            # Enable text widget for updating
            self.output_area.config(state='normal')
            
            # Add user message to history
            self.conversation_history.append(f"Human: {user_input}")
            
            # Display user message
            self.output_area.insert(tk.END, f"You: {user_input}\n\n")
            
            # Create full context from history
            full_context = "\n".join(self.conversation_history)
            
            # Get AI response and clean it
            response = self.llm.invoke(full_context)
            cleaned_response = self.clean_response(response)
            
            # Add AI response to history
            self.conversation_history.append(f"Assistant: {cleaned_response}")
            
            # Display AI response
            self.output_area.insert(tk.END, f"AI: {cleaned_response}\n\n")
            
            # Make text widget read-only again
            self.output_area.config(state='disabled')
            
            # Clear input field
            self.input_field.delete(0, tk.END)
            
            # Scroll to bottom
            self.output_area.see(tk.END)

    def clean_response(self, response):
        """Remove thinking process tags from the response"""
        import re
        return re.sub(r'<think>.*?</think>', '', response, flags=re.DOTALL).strip()

    def attach_pdf(self):
        file_path = filedialog.askopenfilename(
            filetypes=[("PDF files", "*.pdf")]
        )
        if file_path:
            # Read PDF content using the handler
            pdf_text, error = self.pdf_handler.read_pdf(file_path)
            
            if error:
                self.file_label.config(text=f"Error reading PDF: {error}")
                return
                
            # Update file label
            filename = file_path.split('/')[-1]
            self.file_label.config(text=f"Selected: {filename}")
            
            # Enable text widget for updating
            self.output_area.config(state='normal')
            
            # Add PDF content to conversation
            self.conversation_history.append(f"Human: I'm sharing a PDF with the following content:\n{pdf_text}")
            self.output_area.insert(tk.END, f"[PDF uploaded: {filename}]\n\n")
            
            # Get AI response about the PDF
            # full_context = "\n".join(self.conversation_history)
            # response = self.llm.invoke(full_context)
            # cleaned_response = self.clean_response(response)
            
            # Add AI response to history
            self.conversation_history.append(f"Assistant: {pdf_text}")
            
            # Display AI response
            self.output_area.insert(tk.END, f"AI: {pdf_text}\n\n")
            
            # Make text widget read-only again
            self.output_area.config(state='disabled')
            
            # Scroll to bottom
            self.output_area.see(tk.END)

    def attach_proposal(self):
        # Implementation for attaching a proposal PDF
        pass

    def get_ai_response(self, message):
        # Implementation for getting AI response
        pass

    def add_to_chat(self, role, message):
        # Implementation for adding to chat history
        pass