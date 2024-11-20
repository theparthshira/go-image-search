import easyocr

reader = easyocr.Reader(['en'])

text_result = reader.readtext('image.png')

for text in text_result:
    print("result", type(text), text[1], '\n')
