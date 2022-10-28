import React from "react";
import axios from "axios";
import { useSearchParams } from 'react-router-dom';
import '../assets/card-success.css';

let baseUrl = "http://localhost:8098/api/orders?orderNo=";
let postUrl = "http://localhost:8098/api/confirm/";
function Orders() {

   
    const [searchParams, setSearchParams] = useSearchParams();
    var orderNo = searchParams.get("orderNo")
   
    console.log("orderNo" + orderNo)

    const [order, setOrder] = React.useState(null);
    const [paySuccess, setPaySuccess] = React.useState(false);
    const [error, setError] = React.useState(null);

    let urlAPI = baseUrl + orderNo;


    React.useEffect(() => {
        axios
            .get(urlAPI)
            .then((response) => {

                setOrder(response.data.data)
            })
            .catch(error => {
                setError(error)
            })
    }, []);

    function confirm() {
        var data = order
        axios
            .post(postUrl, data)

            .then((response) => {
                setPaySuccess(true)
            })
            .catch(error => {
                setPaySuccess(false)
            })
    }

    if (error) return `Error: ${error.message}`;

    return (


        <div class="container">
            <div class="col-xs-8 col-xs-offset-2 jumbotron text-center">
                {order != null && paySuccess == false ?

                    <>
                        <h1>Your Order</h1>
                        <div class="row justify-content-md-center">
                            <img src="https://gray-kwch-prod.cdn.arcpublishing.com/resizer/34XlWhykDU7BCwTdMhjoyUhKcIY=/980x0/smart/filters:quality(70)/cloudfront-us-east-1.images.arcpublishing.com/gray/5A4FDMHJQVDCDDCUUUMVFQEZT4.png" width={300} height={300}/>
                        </div>
                        <div class="row">
                            <div class="col text-left">
                                <p>ORDER NO:</p>
                            </div>
                            <div class="col text-right">
                                <p>#{order.order_no}</p>
                            </div>
                        </div>
                        <div class="row">
                            <div class="col text-left">
                                <p>MERCHANT:</p>
                            </div>
                            <div class="col text-right">
                                <p>{order.merchant_id}</p>
                            </div>
                        </div>
                        <div class="row">
                            <div class="col text-left">
                                <p>APP ID:</p>
                            </div>
                            <div class="col text-right">
                                <p>{order.app_id}</p>
                            </div>
                        </div>
                        <div class="row">
                            <div class="col text-left">
                                <p>AMOUNT:</p>
                            </div>
                            <div class="col text-right">
                                <p>{order.amount} VND</p>
                            </div>
                        </div>
                        <div class="row">
                            <div class="col text-left">
                                <p>PRODUCT CODE:</p>
                            </div>
                            <div class="col text-right">
                                <p>{order.product_code}</p>
                            </div>
                        </div>
                        <div class="row">
                            <div class="col text-left">
                                <p>DESCRIPTION:</p>
                            </div>
                            <div class="col text-right">
                                <p>{order.description}</p>
                            </div>
                        </div>
                    

                        <div class="row" onClick={confirm}> 
                            <p className="btn btn-primary btn-lg btn-login btn-block">Confirm</p>

                        </div>
                    </>
                    : null}

                {paySuccess ?
                    <>

                        <div class="card">
                            <div class="checkContainer">
                                <i class="checkmark">✓</i>
                            </div>
                            <h1>Success</h1>
                            <p>We received your purchase request;<br /> we'll be in touch shortly!</p>
                        </div>
                    </>

                    : null}
            </div>


        </div>



    );
}

export default Orders;