import React, { useState, useEffect, useMemo } from "react";
import axios from 'axios'

import { Header } from "../components";
import OrderTable from "../components/Table/Table";
import {SelectColumnFilter} from '../components/Table/Filter'

const ordersAPI = "http://localhost:8099/api/orders";

const Status = ({ value }) => {
  return (
      <>
          {value === 0 ?

              (<span className="status-success">
                  Success
              </span>
              )
              : (<span className="status-failed">
                  Failed
              </span>
              )
          }
      </>
  );
};

const Orders = () => {

  const [data, setData] = useState([]);

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


  const columns = useMemo(
    () => [

        {
            Header: "Order No",
            accessor: "orderNo"
        },
        {
            Header: "Merchant",
            accessor: "merchantID",
            
        },
        {
            Header: "App ID",
            accessor: "appID"
        },
        {
            id: "amount",
            Header: "Amount",
            accessor: "amount"
        },
        {
            Header: "Status",
            accessor: "status",
            // disable the filter for particular column
            disableFilters: true,
            Cell: ({ cell: { value } }) => <Status value={value} />
        },
        {
            Header: "Product Code",
            accessor: "productCode",
            Filter: SelectColumnFilter,
            filter: "includes",
        },
        {
            Header: "Description",
            accessor: "description"
        },
        {
            Header: "Create time",
            accessor: "CreateTime"
            
        },

    ],
    []
);


  return (
    <div className="m-2 md:m-10 mt-24 p-2 md:p-10 bg-white rounded-3xl">
      <Header category="Page" title="Orders" />
      <OrderTable data={data} columns={columns} />
    </div>
  );
};
export default Orders;
