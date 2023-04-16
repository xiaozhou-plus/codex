## 动态内存分配

### malloc
    
    ```c
    void *malloc(size_t size);
    ```
malloc的参数就是需要分配的内存字节数，返回值是一个指针，指向分配的内存空间。如果分配失败，返回NULL。

### calloc

    void *calloc(size_t nmemb, size_t size);
    
calloc的参数是需要分配的内存块的个数和每个内存块的大小，返回值是一个指针，指向分配的内存空间。如果分配失败，返回NULL。在返回指向内存的指针之前把它初始化为0。

### realloc
    
    ```c
    void *realloc(void *ptr, size_t size);
    ```