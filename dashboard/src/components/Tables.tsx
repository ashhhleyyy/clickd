import React from "react";
import useSWR from "swr";
import { fetcher } from "../fetcher";
import './tables.css';

interface Ranking {
    key: string
    rank: number
    value: number
}

type KeyDisplayComponent = React.FC<{ s: string }>;

function RankingTable(endpoint: string, title: string, counter: string, KeyDisplay: KeyDisplayComponent = ({ s }) => <>{s}</>): React.FC {
    return () => {
        const { data, error } = useSWR<{ rankings: Ranking[] }>(endpoint, fetcher);

        if (error) return <>{error.toString()}</>;
        if (!data) return <>Loading...</>;

        return <table className="table">
            <thead>
                <tr>
                    <td className="heading">{title}</td>
                    <td className="countHeading">{counter}</td>
                </tr>
            </thead>
            <tbody>
                {data.rankings.map(ranking => <tr key={ranking.key}>
                    <td>
                        <KeyDisplay s={ranking.key} />
                    </td>
                    <td className="count">{ranking.value}</td>
                </tr>)}
            </tbody>
        </table>;
    }
}

export const TopPagesTable = RankingTable('/page-rankings', 'Path', 'Views')

export const TopReferersTable = RankingTable('/referer-rankings', 'Referer', 'Referals', ({ s }) => {
    const url = new URL(s);
    return <span className="icon-item">
        <img src={`https://icons.duckduckgo.com/ip3/${url.host}.ico`} width={24} height={24} className="referer-icon" />
        {s}
    </span>
})
