import React, { useState, useMemo } from 'react'
import { useTable, useFilters, useSortBy } from 'react-table';
import '../assets/table-style.css';

const Status = ({ value }) => {
    return (
        <>
            {value == 0 ?

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

function Table({ data }) {

    const columns = useMemo(
        () => [

            {
                Header: "Order No",
                accessor: "orderNo"
            },
            {
                Header: "Merchant",
                accessor: "merchantID"
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
                Cell: ({ cell: { value } }) => <Status value={value} />
            },
            {
                Header: "Product Code",
                accessor: "productCode"
            },
            {
                Header: "Description",
                accessor: "description"
            },
            {
                Header: "Create time",
                accessor: "CreateTime"
            },

            // {
            //     Header: "Type",
            //     accessor: "show.type"
            // }

            // ,
            // {
            //     Header: "Language",
            //     accessor: "show.language"
            // },
            // {
            //     Header: "Genre(s)",
            //     accessor: "show.genres",
            //     Cell: ({ cell: { value } }) => <Genres values={value} />
            // },
            // {
            //     Header: "Runtime",
            //     accessor: "show.runtime",
            //     Cell: ({ cell: { value } }) => {
            //         const hour = Math.floor(value / 60);
            //         const min = Math.floor(value % 60);
            //         return (
            //             <>
            //                 {hour > 0 ? `${hour} hr${hour > 1 ? "s" : ""} ` : ""}
            //                 {min > 0 ? `${min} min${min > 1 ? "s" : ""}` : ""}
            //             </>
            //         );
            //     }
            // },
            // {
            //     Header: "Status",
            //     accessor: "show.status"
            // }

        ],
        []
    );

    const [filterInput, setFilterInput] = useState("");
    // Use the state and functions returned from useTable to build your UI
    const {
        getTableProps,
        getTableBodyProps,
        headerGroups,
        rows,
        prepareRow,
        setFilter
    } = useTable(
        {
            columns,
            data
        },
        useFilters,
        useSortBy
    );

    const handleFilterChange = e => {
        const value = e.target.value || undefined;
        setFilter("merchantID", value);
        setFilterInput(value);
    };

    // Render the UI for your table
    return (
        <>
            <input
                value={filterInput}
                onChange={handleFilterChange}
                placeholder={"Search Merchant"}
            />
            <table {...getTableProps()}>
                <thead>
                    {headerGroups.map(headerGroup => (
                        <tr {...headerGroup.getHeaderGroupProps()}>
                            {headerGroup.headers.map(column => (
                                <th
                                    {...column.getHeaderProps(column.getSortByToggleProps())}
                                    className={
                                        column.isSorted
                                            ? column.isSortedDesc
                                                ? "sort-desc"
                                                : "sort-asc"
                                            : ""
                                    }
                                >
                                    {column.render("Header")}
                                </th>
                            ))}
                        </tr>
                    ))}
                </thead>
                <tbody {...getTableBodyProps()}>
                    {rows.map((row, i) => {
                        prepareRow(row);
                        return (
                            <tr {...row.getRowProps()}>
                                {row.cells.map(cell => {
                                    return (
                                        <td {...cell.getCellProps()}>{cell.render("Cell")}</td>
                                    );
                                })}
                            </tr>
                        );
                    })}
                </tbody>
            </table>
        </>
    );
}

export default Table