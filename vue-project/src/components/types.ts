interface IQueriesURL {
    allData: boolean;
    pagination: boolean;
    ById: boolean;
}

export interface Clientes {
    customer_id : string,
    full_name : string,
    birth_date : string,
    customer_address : string,
    customer_sector: string,
    customer_postal_code: string,
    phone: number,
    email: string,
    discharge_date: string,
    customer_group: string
}

interface DataPagination {
    page: number;
    amount:number;
}

interface DataById {
    id: number
}

interface IDataObjectQueries {
    pagination : Required<DataPagination>;
    byid: DataById
}

export interface ResultPaginationFetch {
    data_pagination: Clientes[];
    total_rows: number;
}

export interface IStatePagination {
    page: number;
    pages_total: number;
    numbers : [] | [number[], (string[] | null), number[]],
    amount: number,
    slice_page: number;
}

export type ObjectDataQuery = Partial<IDataObjectQueries>
   
export type QueriesUrl = Partial<IQueriesURL>