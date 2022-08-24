# USO: llamar a esta funcion con un argumento en segundos
# python generador.py 2 (genera una medida cada dos segundos)
# python generador.py -poisson 2 (genera medidas con distr. de poisson de media 2 segundos)
# ---
# USAGE: call this program with an argument (in seconds)
# python generador.py 2 (generate a measurement every 2 seconds)
# python generador.py -poisson 2 (generate measurements with poisson distr. with mean 2 seconds)

import requests, datetime, sys, time, random, numpy, json

opts = [opt for opt in sys.argv[1:] if opt.startswith("-")]
args = [arg for arg in sys.argv[1:] if not arg.startswith("-")]
random.seed(datetime.datetime.now())

headers_json = {"Content-Type":"application/json",
"Accept":"application/json",
"Fiware-Service":"SmartSantander",
"Fiware-ServicePath":"/SS"}

data_json = {"urn":"urn:x-iot:smartsantander:u7jcfa:t519","value":26.51,
"uom":"degreeCelsius","location":{"type":"Point","coordinates":[-3.80885,43.46475]},
"phenomenon":"temperature:ambient",
"timestamp":"2021-07-09T15:25:46.905712+02:00"}

i = 1
if "-poisson" in opts:
	media = float(args[0])
	while True :
		timestamp = datetime.datetime.now()
		data_json["timestamp"] = str(timestamp)
		res = requests.post('http://10.10.150.229:5053/notify', headers=headers_json, data=json.dumps(data_json))
		print("Medida enviada en timestamp", str(timestamp))
		time.sleep(-media*numpy.log(random.random())); #Poisson distribution (exponential)
		i = i + 1
		if i >= 5000: # limit to 5000, can remove
			break
else:
	tiempo = float(args[0])
	while True :
		timestamp = datetime.datetime.now()
		data_json["timestamp"] = str(timestamp)
		res = requests.post('http://10.10.150.229:5053/notify', headers=headers_json, data=json.dumps(data_json))
		print("Medida enviada en timestamp", str(timestamp))
		time.sleep(tiempo)
		i = i + 1
		if i >= 5000:
			break
