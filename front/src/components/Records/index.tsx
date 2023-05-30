import React from 'react';
import { IRecord, Record, RecordProps } from './components/record';

export interface RecordsProps {
    records: RecordProps[];
    maxRecords?: number;
    initialRecord?: number;
    setRecords?: React.Dispatch<React.SetStateAction<IRecord[]>>;
}

export const Records = ({
                            records,
                            maxRecords,
                            initialRecord = 0,
                            setRecords
                        }: RecordsProps) => {
    const recordsToShow = records.slice(initialRecord, maxRecords);
    return (
        <ul className="tw-w-full">
            {recordsToShow &&
                recordsToShow.map((record: RecordProps, index: number) => (
                    <Record
                        key={`record-${index}`}
                        {...record}
                        className={`
              ${index + 1 === recordsToShow.length && 'tw-border-b'}`}
                        setRecords={setRecords}
                    />
                ))}
        </ul>
    );
};

export * from './components/record';
