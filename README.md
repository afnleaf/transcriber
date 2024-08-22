# Transcribing Written/Printed Documents using LLMs

Steps
- get image file or array of image files
- check they are images
- get filename
- load file contents
- gemini api connection
- json to json?
- gemini transcribes the content of the image into a form
- temperature?
- format output
- prompt engineer?
- receive contents
- format output locally
- left, original document
- right file contents
- can edit file
- copy or save as file .md or .txt

Technology

lets learn a new language, done with javascript on the backend...

API_KEY="YOUR_API_KEY"
curl -H 'Content-Type: application/json' \
     -d '{"contents":[
            {"role": "user",
              "parts":[{"text": "Give me five subcategories of jazz?"}]}]}' \
     "https://generativelanguage.googleapis.com/v1/models/gemini-pro:generateContent?key=${API_KEY}"

Go?
