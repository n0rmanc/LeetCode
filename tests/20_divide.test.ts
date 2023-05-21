import { divide } from "../20_divide";



describe('Find the Index of the First Occurrence in a String', () => {
    test.each([
        {
            name: "test 1",
            args: {
                dividend: 10, divisor: 3
            },
            want: 3,
        },
        {
            name: "test 2",
            args: {
                dividend: 7, divisor: -3
            },
            want: -2,
        }
    ])(`$name`, async ({ name, args, want }) => {

        abcde();
        expect(divide(args.dividend, args.divisor)).toBe(want);
    });
});

async function abcde() {
    try{
        console.log("abcde");
        return 1;
    }catch(e){
        console.log("catch");
    }finally{
        console.log("finally");
    }
    return 1;
}
