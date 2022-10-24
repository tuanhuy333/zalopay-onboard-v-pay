import React from 'react'
import Orders from '../components/Order'
import { useParams } from 'react-router';
function OrderPage() {
    const { orderId } = useParams();
    console.log("id"+orderId)

    return (
   
        <Orders />
   
  )
}

export default OrderPage