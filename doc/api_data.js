define({ "api": [
  {
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "optional": false,
            "field": "varname1",
            "description": "<p>No type.</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "varname2",
            "description": "<p>With type.</p>"
          }
        ]
      }
    },
    "type": "",
    "url": "",
    "version": "0.0.0",
    "filename": "./doc/main.js",
    "group": "/home/vladimir/go/src/market_analyzer/doc/main.js",
    "groupTitle": "/home/vladimir/go/src/market_analyzer/doc/main.js",
    "name": ""
  },
  {
    "type": "post",
    "url": "/emitters/list",
    "title": "Return list of emitters",
    "name": "emittersListHandler",
    "group": "emitters",
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 OK\n{\n\t\"emitters\": [\n\t\t\"sber\",\n\t\t\"vtbr\",\n\t\t\"yndx\"\n\t]\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 500 Internal Server Error\n{\n\t\"message\":\"Failed to walk directories\"\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./internal/server/emitters_list_handler.go",
    "groupTitle": "emitters"
  },
  {
    "type": "post",
    "url": "/emitters/load",
    "title": "Uploads emitter data",
    "name": "emittersLoadHandler",
    "group": "emitters",
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 OK\n{}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 500 Internal Server Error\n{\n\t\"message\":\"Failed to open file\"\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./internal/server/emitters_load_handler.go",
    "groupTitle": "emitters"
  },
  {
    "type": "post",
    "url": "/emitters/report",
    "title": "Return report about emitter",
    "name": "emittersReportHandler",
    "group": "emitters",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "emitter",
            "description": "<p>Mandatory emitter name</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "allowedValues": [
              "\"normal\"",
              "\"first\"",
              "\"second\"",
              "\"third\"",
              "\"four\"",
              "\"year\""
            ],
            "optional": false,
            "field": "period_mode",
            "defaultValue": "normal",
            "description": "<p>Optional period mod that indicates for which period the report will be displayed.</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "allowedValues": [
              "\"num_with_percent\"",
              "\"number\"",
              "\"percent\""
            ],
            "optional": false,
            "field": "number_mode",
            "defaultValue": "num_with_percent",
            "description": "<p>Optional number mod that specifies which format to output numbers in.</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "precision",
            "description": "<p>Optional field specifies how many decimal places will be displayed in the report.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n  \"emitter\": \"yndx,\n  \"period_mode\": \"normal\",\n  \"number_mode\": \"num_with_percent\",\n  \"precision\": 3\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 OK\n{\n\t\"headers\": [\n\t\t\"#\",\n\t\t\"2016/4\",\n\t\t\"2017/4\",\n\t\t\"2018/4\",\n\t\t\"2019/4\",\n\t\t\"2020/2\",\n\t\t\"2020/3\",\n\t\t\"2020/4\",\n\t\t\"2021/1\",\n\t\t\"2021/2\",\n\t\t\"2021/3\"\n\t],\n\t\"rows\": [\n\t\t[\n\t\t\t\"market_cap\",\n\t\t\t\"-\",\n\t\t\t\"66.000.000.000\",\n\t\t\t\"66.000.000.000(+0%)\",\n\t\t\t\"66.000.000.000(+0%)\",\n\t\t\t\"-\",\n\t\t\t\"-\",\n\t\t\t\"66.000.000.000(+0%)\",\n\t\t\t\"66.000.000.000(+0%)\",\n\t\t\t\"66.000.000.000(+0%)\",\n\t\t\t\"66.000.000.000(+0%)\"\n\t\t],\n\t\t[\n\t\t\t\"sales\",\n\t\t\t\"-\",\n\t\t\t\"-\",\n\t\t\t\"-\",\n\t\t\t\"-\",\n\t\t\t\"1.600.000.000\",\n\t\t\t\"990.000.000(-38%)\",\n\t\t\t\"3.930.000.000(+297%)\",\n\t\t\t\"-\",\n\t\t\t\"1.970.000.000(-50%)\",\n\t\t\t\"1.290.000.000(-35%)\"\n\t\t],\n\t\t[\n\t\t\t\"earnings\",\n\t\t\t\"-\",\n\t\t\t\"-\",\n\t\t\t\"-\",\n\t\t\t\"-\",\n\t\t\t\"-162.000.000\",\n\t\t\t\"256.000.000(+258%)\",\n\t\t\t\"1.660.000.000(+548%)\",\n\t\t\t\"-\",\n\t\t\t\"-180.000.000(-111%)\",\n\t\t\t\"245.000.000(+236%)\"\n\t\t],\n\t\t[\n\t\t\t\"debts\",\n\t\t\t\"-\",\n\t\t\t\"-\",\n\t\t\t\"1.480.000.000\",\n\t\t\t\"1.320.000.000(-11%)\",\n\t\t\t\"-\",\n\t\t\t\"-\",\n\t\t\t\"1.400.000.000(+6%)\",\n\t\t\t\"-\",\n\t\t\t\"-\",\n\t\t\t\"2.700.000.000(+93%)\"\n\t\t],\n\t\t[\n\t\t\t\"p/e\",\n\t\t\t\"-\",\n\t\t\t\"-\",\n\t\t\t\"-\",\n\t\t\t\"-\",\n\t\t\t\"-\",\n\t\t\t\"-\",\n\t\t\t\"40\",\n\t\t\t\"-\",\n\t\t\t\"-367\",\n\t\t\t\"269\"\n\t\t],\n\t\t[\n\t\t\t\"p/s\",\n\t\t\t\"-\",\n\t\t\t\"-\",\n\t\t\t\"-\",\n\t\t\t\"-\",\n\t\t\t\"-\",\n\t\t\t\"-\",\n\t\t\t\"17\",\n\t\t\t\"-\",\n\t\t\t\"34\",\n\t\t\t\"51\"\n\t\t]\n\t]\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 500 Internal Server Error\n{\n\t\"message\":\"Failed to parse file\"\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./internal/server/emitters_report_handler.go",
    "groupTitle": "emitters"
  }
] });
