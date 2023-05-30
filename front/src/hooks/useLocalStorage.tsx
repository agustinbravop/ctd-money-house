import { useEffect, useRef, useState } from 'react';

export function useLocalStorage(
    key: string,
    { serialize = JSON.stringify, deserialize = JSON.parse } = {}
) {
    const [value, setValue] = useState(() => {
        const valueInLocalStorage = window.localStorage.getItem(key);
        if (valueInLocalStorage) {
            return deserialize(valueInLocalStorage);
        }
        return null;
    });

    const prevKeyRef = useRef(key);

    useEffect(() => {
        const prevKey = prevKeyRef.current;

        if (prevKey !== key) {
            window.localStorage.remove(prevKey);
        }
        prevKeyRef.current = key;
        window.localStorage.setItem(key, serialize(value));
    }, [value, serialize, key]);

    return [value, setValue];
}
