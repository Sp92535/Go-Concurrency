### **Experiment 9: Context for Goroutine Cancellation**  

#### **Understanding Context in Go**  
The **`context` package** in Go provides a way to **control goroutines' lifecycles**, enabling **cancellation, deadlines, and timeouts**.  

🔹 **Why Use It?**  
- Avoids **orphaned goroutines** running indefinitely.  
- Provides **graceful shutdown** of concurrent operations.  
- Useful for **API calls, background tasks, and worker pools**.  

---

### **Key Concepts**  

1. **`context.Background()`**  
   - The root context, used when no parent context exists.  

2. **`context.WithCancel()`**  
   - Creates a new context with cancellation capability.  
   - Calling `cancel()` stops all derived goroutines.  

3. **`context.WithTimeout()`**  
   - Cancels the context **after a set duration**.  

4. **Checking Context Inside a Goroutine**  
   - `ctx.Done()` returns a channel that **closes when canceled**.  
   - Use `select` to **listen for cancellation signals**.  

---

### **Problem Statement**  
**Implement a long-running goroutine that prints numbers continuously. Use** `context.WithTimeout()` **to stop it after 3 seconds.**  

---

### **Steps for Implementation**  

1. **Create a context with a 3-second timeout (`context.WithTimeout`).**  
2. **Launch a goroutine that prints numbers in an infinite loop.**  
3. **Use a `select` statement inside the goroutine to listen for `ctx.Done()`.**  
4. **When `ctx.Done()` is triggered, exit the loop and print a message.**  