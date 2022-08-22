#!flask/bin/python

import waitress, requests
from flask import Flask, request
from flask_restful import Api, Resource

app = Flask(__name__, static_url_path="")
api = Api(app)


class SSmiddleware(Resource):
    def post(self):
      request.get_data()
      headers_dic = {"Content-Type": "application/json",
       "Accept": "application/json",
       "Fiware-Service":"SmartSantander",
       "Fiware-ServicePath":"/SS"}
      res = requests.post('http://localhost:5053/notify', headers=headers_dic, data=request.data)
      pass

api.add_resource(SSmiddleware, '/SSapi', endpoint='SSapi')

if __name__ == '__main__':
    #app.run(debug=True, host='0.0.0.0', port=3000)
    waitress.serve(app, host="0.0.0.0", port=3000)
