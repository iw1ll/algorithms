const strStr = (haystack: string, needle: string): number => {
    if (needle.length === 0) return 0;

    for (let i = 0; i <= haystack.length - needle.length; i++) {
        let j = 0;
        while (j < needle.length && haystack[i + j] === needle[j]) {
            j++;
        }
        if (j === needle.length) {
            return i;
        }
    }
    return -1;
};

const testPerformance = (haystack: string, needle: string) => {
    console.time('Custom Implementation');
    const customResult = strStr(haystack, needle);
    console.timeEnd('Custom Implementation');

    console.time('Built-in Implementation');
    const builtInResult = haystack.indexOf(needle);
    console.timeEnd('Built-in Implementation');

    console.log(`Custom Implementation Result: ${customResult}`);
    console.log(`Built-in Implementation Result: ${builtInResult}`);
    
    if (customResult === builtInResult) {
        console.log("Результаты совпадают!");
    } else {
        console.log("Результаты не совпадают!");
    }
    console.log('-----------------------------------');
};

const generateTestCase = (length: number, needleLength: number): [string, string] => {
    const haystack = Math.random().toString(36).substring(2, length + 2).repeat(Math.ceil(length / 5));
    const needle = Math.random().toString(36).substring(2, needleLength + 2);
    return [haystack, needle];
};

export function startTestPerformance() {
for (let i = 0; i < 5; i++) {
    const [haystack, needle] = generateTestCase(100000, 10); // Тест с длинным haystack (100,000 символов) и коротким needle (10 символов)
    console.log(`Test case ${i + 1}: haystack length = ${haystack.length}, needle = '${needle}'`);
    testPerformance(haystack, needle);
}
}
