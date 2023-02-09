import axios from 'axios'

const apiClient = axios.create({
    baseURL: 'http://127.0.0.1:7123/', //api/v1.0/
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
            return apiClient.put(refName + '/'+ detailItem.id,
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
}