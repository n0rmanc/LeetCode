export function strStr(haystack: string, needle: string): number {
    if (needle.length === 0) return 0;
    if (haystack.length === 0) return -1;
    if (haystack.length < needle.length) return -1;
    for (let i = 0; i < haystack.length; i++) {
        if (haystack[i] === needle[0]) {
            let j = 0;
            while (j < needle.length) {
                if (haystack[i + j] !== needle[j]) break;
                j++;
            }
            if (j === needle.length) return i;
        }
    }
    return -1;
};
