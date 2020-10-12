package org.openapitools.api

import org.openapitools.model.Order
interface StoreApiService {

	fun deleteOrder(orderId: kotlin.String): Unit

	fun getInventory(): 

	fun getOrderById(orderId: kotlin.Long): Order

	fun placeOrder(body: Order): Order
}
