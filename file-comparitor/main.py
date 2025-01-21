# Main entry point for app that will parse text from pdf and docx files and send to LLM

import os
import tkinter as tk
from dotenv import load_dotenv
from ui import ChatApp

load_dotenv()

if __name__ == "__main__":
    root = tk.Tk()
    app = ChatApp(root, model=os.getenv("model"))
    root.mainloop()

