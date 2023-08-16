export default function fib(n = 20) {
    if(n === 1 || n === 2) {
        return 1;
    }
    return fib(n-1) + fib(n-2);
}