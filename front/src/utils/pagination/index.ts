import { PAGE } from '../../constants';

export const pageQuery = (route: string, page: number) =>
    `${route}${page === 1 ? '' : `?${PAGE}${page}`}`;