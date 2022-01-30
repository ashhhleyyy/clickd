import { Fetcher } from "swr";

export const fetcher: Fetcher<any> = (path: string) => fetch(path).then(res => res.json());
