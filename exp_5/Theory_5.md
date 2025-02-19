### **Experiment 5: Select Statement for Multiple Channels**  

#### **Understanding `select` in Go**  
The `select` statement in Go is used to **handle multiple channel operations** simultaneously. It allows a goroutine to wait on multiple channels and proceed with whichever one is ready first.  

#### **Key Concepts**  

1. **Using `select` with Multiple Channels**  
   - `select` listens to multiple channels and executes the first one that receives data.  
   - If multiple channels are ready, **one is chosen randomly**.  
   - If no channels are ready, it **blocks** unless a `default` case is provided.  

   ```go
   select {
   case msg1 := <-ch1:
       fmt.Println("Received from ch1:", msg1)
   case msg2 := <-ch2:
       fmt.Println("Received from ch2:", msg2)
   default:
       fmt.Println("No messages available, moving on...")
   }
   ```

2. **Avoiding Deadlocks**  
   - If no channels are ready and there’s no `default` case, `select` **blocks indefinitely**.  
   - A `time.After` channel can be used for timeouts.  

3. **Using `select` in Real-World Scenarios**  
   - Handling multiple network requests.  
   - Listening to multiple worker goroutines.  
   - Implementing timeouts for operations.  

---

### **Problem Statement**  
Implement a system where **multiple goroutines** send messages to different channels, and the `select` statement in the main function picks whichever message is available first.  

---

### **Steps for Implementation**  

1. **Create two or more channels for communication.**  
2. **Launch goroutines that send messages at different intervals.**  
3. **Use `select` in the main function to listen to all channels.**  
4. **Observe how `select` picks the first available message.**  
5. **Introduce a timeout using `time.After` to prevent indefinite blocking.**  

Try this out and see how Go handles multiple channels efficiently! 🚀