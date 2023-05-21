export function divide(dividend: number, divisor: number): number {
    if (dividend === 0) return 0;
    if (divisor === 1) return dividend;
    if (divisor === -1) {
        if (dividend > -2147483648) return -dividend;
        return 2147483647;
    }
    let isNegative = false;
    if (dividend > 0 && divisor < 0 || dividend < 0 && divisor > 0) isNegative = true;
    dividend = Math.abs(dividend);
    divisor = Math.abs(divisor);
    let quotient = 0;
    while (dividend >= divisor) {
        let temp = divisor;
        let multiple = 1;
        while (dividend >= temp) {
            dividend -= temp;
            quotient += multiple;
            temp += temp;
            multiple += multiple;
        }
    }
    return isNegative ? -quotient : quotient;
};

