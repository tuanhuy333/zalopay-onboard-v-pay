import React, { useState } from 'react'
import { useTable, useFilters, useGlobalFilter, useSortBy, usePagination } from 'react-table';
import { GlobalFilter, DefaultFilterForColumn } from "./Filter";

import '../../assets/table-style.css';

const OrderTable = ({ columns, data }) => {


    // State
    const [filterInput, setFilterInput] = useState("");


    // Use the state and functions returned from useTable to build your UI
    const {
        getTableProps,
        getTableBodyProps,
        setFilter,
        headerGroups,
        page,
        nextPage,
        previousPage,
        canPreviousPage,
        canNextPage,
        pageOptions,
        state,
        visibleColumns,
        gotoPage,
        pageCount,
        setPageSize,
        prepareRow,
        setGlobalFilter,
        preGlobalFilteredRows,


    } = useTable(
        {
            columns,
            data,
            defaultColumn: { Filter: DefaultFilterForColumn },
        },
        useFilters,
        useGlobalFilter,
        useSortBy,
        usePagination
    );


    const { pageIndex, pageSize } = state;


    const handleFilterChange = e => {
        const value = e.target.value || undefined;
        setFilter("merchantID", value);
        setFilterInput(value);
    };

    // Render the UI for your table
    return (
        <>

            <table {...getTableProps()}>
                <thead>
                    <tr>
                        <th
                            colSpan={visibleColumns.length}
                            style={{
                                textAlign: "center",
                            }}
                        >
                            {/* rendering global filter */}
                            <GlobalFilter
                                preGlobalFilteredRows={preGlobalFilteredRows}
                                globalFilter={state.globalFilter}
                                setGlobalFilter={setGlobalFilter}
                            />
                        </th>
                    </tr>
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
                                    {/* Render the columns filter UI */}
                                    <div>{column.canFilter ? column.render("Filter") : null}</div>
                                </th>
                            ))}
                        </tr>
                    ))}
                </thead>
                <tbody {...getTableBodyProps()}>
                    {page.map((row) => {
                        prepareRow(row);
                        return (
                            <tr {...row.getRowProps()}>
                                {row.cells.map((cell) => {
                                    return (
                                        <td {...cell.getCellProps()}>{cell.render("Cell")}</td>
                                    );
                                })}
                            </tr>
                        );
                    })}
                </tbody>

            </table>
            {/* Pagging */}
            <div>
                <button onClick={() => gotoPage(0)} disabled={!canPreviousPage}>
                    {"<<"}
                </button>{" "}
                <button onClick={() => previousPage()} disabled={!canPreviousPage}>
                    Previous
                </button>{" "}
                <button onClick={() => nextPage()} disabled={!canNextPage}>
                    Next
                </button>{" "}
                <button onClick={() => gotoPage(pageCount - 1)} disabled={!canNextPage}>
                    {">>"}
                </button>{" "}
                <span>
                    Page{" "}
                    <strong>
                        {pageIndex + 1} of {pageOptions.length}
                    </strong>{" "}
                </span>
                <span>
                    | Go to page:{" "}
                    <input
                        type="number"
                        defaultValue={pageIndex + 1}
                        onChange={(e) => {
                            const pageNumber = e.target.value
                                ? Number(e.target.value) - 1
                                : 0;
                            gotoPage(pageNumber);
                        }}
                        style={{ width: "50px" }}
                    />
                </span>{" "}
                <select
                    value={pageSize}
                    onChange={(e) => setPageSize(Number(e.target.value))}
                >
                    {[5, 10, 15].map((pageSize) => (
                        <option key={pageSize} value={pageSize}>
                            Show {pageSize}
                        </option>
                    ))}
                </select>
            </div>
        </>


    );
}

export default OrderTable;