import tkinter as tk
from tkinter import ttk, filedialog
from langchain_community.llms import Ollama
from pdf_handler import PDFHandler

class ChatApp:
    def __init__(self, root, model):
        self.root = root
        self.root.title("AI Chat Interface")
        
        # Initialize Ollama and PDF handler
        self.llm = Ollama(model=model)
        self.pdf_handler = PDFHandler()
        
        # Initialize conversation history
        self.conversation_history = []
        
        # Create main container
        main_frame = ttk.Frame(root, padding="10")
        main_frame.grid(row=0, column=0, sticky=(tk.W, tk.E, tk.N, tk.S))
        
        # Create and configure text display area
        self.output_area = tk.Text(main_frame, height=30, width=50, wrap=tk.WORD)
        self.output_area.grid(row=0, column=0, columnspan=2, pady=5)
        
        # Add scrollbar
        scrollbar = ttk.Scrollbar(main_frame, orient='vertical', command=self.output_area.yview)
        scrollbar.grid(row=0, column=2, sticky='ns')
        self.output_area['yscrollcommand'] = scrollbar.set
        
        # Make the text read-only
        self.output_area.config(state='disabled')
        
        # Create input field
        self.input_field = ttk.Entry(main_frame, width=40)
        self.input_field.grid(row=1, column=0, pady=5)

        # Create file label
        self.file_label = ttk.Label(main_frame, text="No file selected")
        self.file_label.grid(row=2, column=0, pady=5, sticky='w')

        # add button to attach pdf
        self.attach_pdf_button = ttk.Button(main_frame, text="Attach PDF", command=self.attach_pdf)
        self.attach_pdf_button.grid(row=1, column=1, pady=5, padx=5)
        
        # Create send button
        send_button = ttk.Button(main_frame, text="Send", command=self.send_message)
        send_button.grid(row=1, column=2, pady=5, padx=5)
        
        # Bind Enter key to send message
        self.input_field.bind("<Return>", lambda e: self.send_message())

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
            full_context = "\n".join(self.conversation_history)
            response = self.llm.invoke(full_context)
            cleaned_response = self.clean_response(response)
            
            # Add AI response to history
            self.conversation_history.append(f"Assistant: {cleaned_response}")
            
            # Display AI response
            self.output_area.insert(tk.END, f"AI: {cleaned_response}\n\n")
            
            # Make text widget read-only again
            self.output_area.config(state='disabled')
            
            # Scroll to bottom
            self.output_area.see(tk.END)