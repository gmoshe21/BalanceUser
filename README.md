# BalanceUser  

Запуск go run cmd/main.go  
В файле config/config.json конфигурации сервера
Имеется 4 запроса :  
    POST /control/debiting_money  
        принимает json в теле вида :  
                        {  
                            "ID": "1",  
                            "Money": "243"  
                        }  
        списывает денежные средства со счета,

    POST /control/crediting_money  
        принимает json в теле вида :  
                        {  
                            "ID": "2",  
                            "Money": "112"  
                        }  
        добавляет денежные средства со счета,

    POST /control/transfer  
        принимает json в теле вида :  
                        {  
                            "Sender": "1",
                            "Recipient": "2",  
                            "Money": "50"  
                        }  
        перечисление денег,

    GET /control/get_balance/Id=1
        выдает баланс пользователя,


