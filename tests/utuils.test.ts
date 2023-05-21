import { titleize } from "../utils";


describe('titleize', () => {
    it('should titleize a string with underscores', () => {
        const str = "__this__is__a__test";
        const titleizedStr = titleize(str);
        expect(titleizedStr).toBe("This Is A Test");
    });

    it('should titleize a string without underscores', () => {
        const str = "the quick brown fox jumps over the lazy dog";
        const titleizedStr = titleize(str);
        expect(titleizedStr).toBe("The Quick Brown Fox Jumps Over The Lazy Dog");
    });

    it('should titleize a string with mixed case and underscores', () => {
        const str = "tHE_quICK_bROwN_fOX_jUMPs_oVER_tHE_LAZY_dOG";
        const titleizedStr = titleize(str);
        expect(titleizedStr).toBe("The Quick Brown Fox Jumps Over The Lazy Dog");
    });

    it('should titleize a string with leading/trailing underscores', () => {
        const str = "__ the quick brown fox jumps over the lazy dog __";
        const titleizedStr = titleize(str);
        expect(titleizedStr).toBe("The Quick Brown Fox Jumps Over The Lazy Dog");
    });

    it('should titleize a string with multiple consecutive underscores', () => {
        const str = "the___quick__brown_fox__jumps__over_the_lazy_dog";
        const titleizedStr = titleize(str);
        expect(titleizedStr).toBe("The Quick Brown Fox Jumps Over The Lazy Dog");
    });
});
