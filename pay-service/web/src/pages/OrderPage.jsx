import React from 'react'
import Posts from '../components/Order'
import { useParams } from 'react-router';
function OrderPage() {
    const { orderId } = useParams();
    console.log("id"+orderId)

    return (
   
        <Posts />
   
  )
}

export default OrderPage