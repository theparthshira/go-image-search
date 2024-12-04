from confluent_kafka import Consumer, KafkaException
import easyocr
import requests
from base64 import b64encode
from PIL import Image
import io

def basic_auth(username, password):
    token = b64encode(f"{username}:{password}".encode('utf-8')).decode("ascii")
    return f'Basic {token}'

def get_image(file_name):
    print('getting image')
    
    header = {"Authorization": basic_auth("5640111a3b80","00397a4a6a46c8de77b793cf0853a0ceb18a62a29a")}
    auth_response = requests.get("https://api.backblazeb2.com/b2api/v3/b2_authorize_account", headers=header)

    auth_data = auth_response.json()

    print("auth_data['authorizationToken']", auth_data['authorizationToken'])

    return f"https://f003.backblazeb2.com/file/go-image-search/{file_name}?Authorization={auth_data['authorizationToken']}"

def image_ocr(image):
    print("found image", image)
    # image_show = Image.open(io.BytesIO(image))
    # image_show.show()
    
    reader = easyocr.Reader(['en'])
    text_result = reader.readtext(image)

    for text in text_result:
        print("result", type(text), text[1], '\n')


def main():
    print("Kafka started...")
    
    consumer = Consumer({
        'bootstrap.servers': 'localhost:9092',
        'group.id': 'example-group',
        'auto.offset.reset': 'earliest'
    })

    topic = 'comments'
    consumer.subscribe([topic])

    try:
        while True:
            msg = consumer.poll(1.0)  # Poll messages from Kafka
            if msg is None:
                continue
            if msg.error():                
                print(f"Error: {msg.error()}")
            else:
                # file_name = get_image(msg.value().decode('utf-8'))
                image_text = image_ocr(msg.value().decode('utf-8'))
    except KeyboardInterrupt:
        pass
    finally:
        consumer.close()

if __name__ == "__main__":
    main()
