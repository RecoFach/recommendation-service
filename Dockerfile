FROM python:3.8-slim

ADD recommendation.py .
ADD subjects.csv .
ADD requirements.txt .

WORKDIR .

RUN pip3 install --no-cache-dir -r requirements.txt

EXPOSE 2000

CMD ["gunicorn", "-w", "4","-b","0.0.0.0:2000", "recommendation:app" ]
