import { Area, CartesianGrid, ComposedChart, ResponsiveContainer, Tooltip, XAxis, YAxis } from "recharts";
import useSWR from "swr";
import { fetcher } from "../fetcher";

interface ViewsByDayData {
    date: number
    count: number
}  

export function ViewsByTime() {
    const { data, error } = useSWR<{ values: ViewsByDayData[] }>('/views-by-day', fetcher);

    if (error) return <>{error.toString()}</>;
    if (!data) return <>Loading...</>;

    return <ResponsiveContainer width="100%" height={250}>
        <ComposedChart width={730} height={250} data={data.values} margin={{ top: 10, right: 30 }}>
            <defs>
                <linearGradient id="grad" x1="0" y1="0" x2="0" y2="1">
                    <stop offset="5%" stopColor="#8884d8" stopOpacity={0.8}/>
                    <stop offset="95%" stopColor="#8884d8" stopOpacity={0}/>
                </linearGradient>
            </defs>
            <XAxis dataKey="date" type="number" domain={['auto', 'auto']} name='Time' tickFormatter={(time: number) => formatDate(time * 1000)} />
            <YAxis />
            <CartesianGrid strokeDasharray="3 3" />
            <Tooltip />
            <Area type="monotone" dataKey="count" stroke="#8884d8" fillOpacity={1} fill="url(#grad)" />
        </ComposedChart>
    </ResponsiveContainer>
}

function formatDate(timestamp: number): string {
    const date = new Date(timestamp);
    return `${date.getDate()}/${date.getMonth() + 1}/${date.getFullYear()}`;
}
