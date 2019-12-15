#include <stdio.h>
#include "libgoutil.h"

int main(void) {
    int a = 12;
    int b = 99;

    printf("goutil.Add(%d, %d)=%d\n", a, b, (int)Add(a,b));
    printf("goutil.Cosine(1)=%f\n", (float)Cosine(1.0));

    GoInt data[6] = { 77, 12, 5, 99, 28, 23};
    GoSlice nums = {data, 6, 6};
    Sort(nums);
    for (int i=0; i<6; i++) {
        printf("%d,", (int)((GoInt *)nums.data)[i]);
    }
    printf("\nsizeof(data)=%lu\n", sizeof(data));

    GoString msg = { "Hello from C!", 13};
    Log(msg);
    return 0;
}
