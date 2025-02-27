### **Experiment 10: Parallelism with GOMAXPROCS**  

#### **Understanding Parallelism in Go**  
Go provides **runtime.GOMAXPROCS(n)** to control how many OS threads execute Go code simultaneously. By default, it is set to the number of logical CPU cores.  

Parallelism differs from concurrency:  
- **Concurrency** = Task switching among multiple goroutines.  
- **Parallelism** = Actual simultaneous execution of goroutines on multiple cores.  

#### **Key Concepts**  

1. **GOMAXPROCS and CPU Cores**  
   - Increasing `GOMAXPROCS` allows Go to use more CPU cores for parallel execution.  
   - Setting it too high can cause overhead, while too low underutilizes resources.  

2. **Measuring Performance Impact**  
   - We compare execution times by varying `GOMAXPROCS` values.  
   - CPU-intensive tasks like matrix multiplication or factorial computation highlight differences.  

3. **Goroutines and Parallel Execution**  
   - With `GOMAXPROCS = 1`, goroutines run sequentially on a single core.  
   - With `GOMAXPROCS = NumCPU()`, tasks are executed in parallel on multiple cores.  

---

### **Problem Statement**  
Write a program that performs a CPU-intensive task (e.g., summing large numbers) using multiple goroutines.  
- **Measure execution time for different values of `GOMAXPROCS`**.  
- **Compare results and determine optimal parallelism for your system**.  

---

### **Steps for Implementation**  

1. **Set `runtime.GOMAXPROCS(n)` to different values and observe execution time.**  
2. **Run multiple goroutines performing CPU-heavy operations.**  
3. **Use `sync.WaitGroup` to ensure all tasks complete before measuring execution time.**  
4. **Compare performance with different CPU core utilizations.**  