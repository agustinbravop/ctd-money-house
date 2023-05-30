import { useEffect } from 'react';
import { useNavigate, useSearchParams } from 'react-router-dom';
import { IRecord } from '../components/Records/';
import { ROUTES } from '../constants';

export const usePagination = (
    records: IRecord[],
    recordsPerPage: number,
    initialPage = 1
) => {
    const [searchParams] = useSearchParams();
    const page = searchParams.get('page');
    const pageNumber = (page && parseInt(page, 10)) || initialPage;
    const numberOfPages = Math.ceil(records.length / recordsPerPage);
    const navigate = useNavigate();

    useEffect(() => {
        if (numberOfPages !== 0 && pageNumber > numberOfPages) {
            navigate(ROUTES.NOT_FOUND);
        }
    }, [pageNumber, numberOfPages, navigate]);

    const isRecordsGreeterThanOnePage = records.length > recordsPerPage;
    return {
        pageNumber,
        numberOfPages,
        isRecordsGreeterThanOnePage,
    };
};
