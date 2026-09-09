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