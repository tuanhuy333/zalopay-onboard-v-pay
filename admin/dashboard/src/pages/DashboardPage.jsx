import React, { useState, useEffect, useMemo } from 'react'
import Table from '../components/Table'
import axios from 'axios';


const ordersAPI = "http://localhost:8099/api/orders";

function AdminPage() {


    const [data, setData] = useState([])

    // function getColumns() {
    //     if (columns) {
    //     return columns.map(key => {
    //       return {
    //         id: key.toString(),
    //         Header: key.toString(),
    //         accessor: key.toString()
    //       };
    //     });
    //   }}



    const getData = () => {
        axios.get(ordersAPI)
            .then(res => {
                console.log(res.data)
                setData(res.data.data)
            })
            .catch(error => {
                console.log(error)
            })
    }



    useEffect(() => {
        getData()
    }, []);



    return (

        <Table data={data} />
    )

}

export default AdminPage