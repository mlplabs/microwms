import axios from 'axios'

const apiClient = axios.create({
    baseURL: 'http://microwms.ru/api/v1.0/', //api/v1.0/
    withCredentials: false,
    headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
    }
})

export default {
    StoreItemReference(refName, detailItem){
        if (detailItem.id === 0) {
            return apiClient.post(refName,
                detailItem,
                {
                    headers: {
                        'Content-type': 'application/x-www-form-urlencoded'
                    }
                }
            )
        }else {
            return apiClient.put(refName,
                detailItem,
                {
                    headers: {
                        'Content-type': 'application/x-www-form-urlencoded'
                    }
                }
            )
        }
    },

    GetItemsReference(refName, page, limit, offset){
        return apiClient.get(`${refName}?l=${limit}&o=${offset}`)
    },

    GetItemReference(refName, id){
        return apiClient.get(`${refName}/${id}`)
    },
    PrintItemReference(refName, id){
        return apiClient.get(`print/${refName}/${id}`)
    },
    GetSuggestionReference(refName, text){
        return apiClient.get(`suggestion/${refName}/${text}`)
    },

    DeleteItemReference(refName, id){
        return apiClient.delete(`${refName}/${id}`)
    },

    SearchFirmByInn(text) {
        return apiClient.get('firms/find/inn/'+text)
    },
    GetHwPrinters(){
        return apiClient.get(`printers`)
    },
    GetReceiptDoc(refName, id){
        return apiClient.get(`${refName}/${id}`)
    },
    GetReceiptDocs(refName, page, limit, offset){
        return apiClient.get(`${refName}?l=${limit}&o=${offset}`)
    },
    StoreReceiptDoc(refName, detailItem){
        if (detailItem.id === 0) {
            return apiClient.post(refName,
                detailItem,
                {
                    headers: {
                        'Content-type': 'application/x-www-form-urlencoded'
                    }
                }
            )
        }else {
            return apiClient.put(refName,
                detailItem,
                {
                    headers: {
                        'Content-type': 'application/x-www-form-urlencoded'
                    }
                }
            )
        }
    },

    DeleteReceiptDoc(refName, id){
        return apiClient.delete(`${refName}/${id}`)
    },
    GetEnum(enumName){
        return apiClient.get(`enum/${enumName}`)
    },
    GetReport(reportName){
        return apiClient.get(`reports/${reportName}`)
    },

    ErrorProcessing(error){
        console.log(error)
        if (error.response) {
            // client received an error response (5xx, 4xx)
            if (error.response.status === 404){
                this.statusText = "Ничего не найдено ("
            }else {
                this.statusText = "Произошла ошибка ("+error.response.status+")"
            }
        } else if (error.request) {
            // client never received a response, or request never left
        } else {
            // anything else
        }

    }
}