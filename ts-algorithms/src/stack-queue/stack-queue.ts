// stack --> last in first out --> array, linked-list(-->)
// queue --> first in first out --> linked-list(<-->)

export function isValid(s: string): boolean {
    const pair: Record<string, string> = {
        '(': ')',
        '[': ']',
        '{': '}'
    };
    const stack: string[] = [];

    for (const char of s) {
        if (char in pair) {
            stack.push(char);
        } else {
            if (stack.length === 0) {
                return false;
            }

            const prev = stack.pop();

            if (prev === undefined) {
                return false;
            }

            if (char !== pair[prev]) {
                return false;
            }
        }
    }

    return stack.length === 0;
}

export function removeDuplicatesS(s: string): string {
    const stack: string[] = [];

    for (const char of s) {
        if (stack.length > 0 && stack[stack.length - 1] === char){
            stack.pop();
        } else {
            stack.push(char);
        }
    }
    
    return stack.join('');
};