# Go Beyond: Executable Code

This repo contains executable Go programs from the [Go Beyond: Mastering the Art of Go Programming: A Comprehensive Practical Guide to Golang Development](https://www.amazon.com/dp/B0FT1K7KSJ) book.

## Programs by Chapter


### chapter11.md

- **The Problem Without DI** ([ch11_the_problem_without_di](./go-beyond-example-code/ch11_the_problem_without_di/)) - Line 36
- **The Solution With DI** ([ch11_the_solution_with_di](./go-beyond-example-code/ch11_the_solution_with_di/)) - Line 88
- **Basic Constructor Injection** ([ch11_basic_constructor_injection](./go-beyond-example-code/ch11_basic_constructor_injection/)) - Line 162
- **Multiple Dependencies** ([ch11_multiple_dependencies](./go-beyond-example-code/ch11_multiple_dependencies/)) - Line 248
- **Service Layer with Interfaces** ([ch11_service_layer_with_interfaces](./go-beyond-example-code/ch11_service_layer_with_interfaces/)) - Line 351
- **Mock Dependencies** ([ch11_mock_dependencies](./go-beyond-example-code/ch11_mock_dependencies/)) - Line 515
- **Simple DI Container** ([ch11_simple_di_container](./go-beyond-example-code/ch11_simple_di_container/)) - Line 679
- **Advanced DI Container with Factory Functions** ([ch11_advanced_di_container_with_factory_functions](./go-beyond-example-code/ch11_advanced_di_container_with_factory_functions/)) - Line 771
- **Example** ([ch11_example](./go-beyond-example-code/ch11_example/)) - Line 863

### chapter12.md

- **Example** ([ch12_example](./go-beyond-example-code/ch12_example/)) - Line 53
- **example shows a basic mock that maintains state and can be configured to return different responses** ([ch12_example_shows_a_basic_mock_that_maintains_state_an](./go-beyond-example-code/ch12_example_shows_a_basic_mock_that_maintains_state_an/)) - Line 167
- **example shows how to track method calls and their parameters** ([ch12_example_shows_how_to_track_method_calls_and_their_](./go-beyond-example-code/ch12_example_shows_how_to_track_method_calls_and_their_/)) - Line 332
- **example shows how to test a user service with various scenarios** ([ch12_example_shows_how_to_test_a_user_service_with_vari](./go-beyond-example-code/ch12_example_shows_how_to_test_a_user_service_with_vari/)) - Line 566
- **example demonstrates a sophisticated cache mock that tracks detailed behavior and provides comprehensive verification** ([ch12_example_demonstrates_a_sophisticated_cache_mock_th](./go-beyond-example-code/ch12_example_demonstrates_a_sophisticated_cache_mock_th/)) - Line 1007

### chapter13.md

- **Example** | Web server handling multiple requests | Video rendering using all CPU cores |** ([ch13_example_web_server_handling_multiple_requests_vide](./go-beyond-example-code/ch13_example_web_server_handling_multiple_requests_vide/)) - Line 45
- **Basic Goroutines** ([ch13_basic_goroutines](./go-beyond-example-code/ch13_basic_goroutines/)) - Line 158
- **Goroutine Lifecycle and Management** ([ch13_goroutine_lifecycle_and_management](./go-beyond-example-code/ch13_goroutine_lifecycle_and_management/)) - Line 201
- **Goroutine Error Handling** ([ch13_goroutine_error_handling](./go-beyond-example-code/ch13_goroutine_error_handling/)) - Line 248
- **Goroutine Pool Pattern** ([ch13_goroutine_pool_pattern](./go-beyond-example-code/ch13_goroutine_pool_pattern/)) - Line 295
- **Basic Channels** ([ch13_basic_channels](./go-beyond-example-code/ch13_basic_channels/)) - Line 392
- **Advanced Channel Patterns** ([ch13_advanced_channel_patterns](./go-beyond-example-code/ch13_advanced_channel_patterns/)) - Line 428
- **main_program** ([ch13_main_program](./go-beyond-example-code/ch13_main_program/)) - Line 482
- **main_program** ([ch13_main_program_975267](./go-beyond-example-code/ch13_main_program_975267/)) - Line 553
- **Buffered vs Unbuffered Channels** ([ch13_buffered_vs_unbuffered_channels](./go-beyond-example-code/ch13_buffered_vs_unbuffered_channels/)) - Line 620
- **Channel Directions** ([ch13_channel_directions](./go-beyond-example-code/ch13_channel_directions/)) - Line 661
- **WaitGroup: Waiting for Goroutines** ([ch13_waitgroup_waiting_for_goroutines](./go-beyond-example-code/ch13_waitgroup_waiting_for_goroutines/)) - Line 706
- **Mutex: Protecting Shared Data** ([ch13_mutex_protecting_shared_data](./go-beyond-example-code/ch13_mutex_protecting_shared_data/)) - Line 740
- **RWMutex: Read-Write Locks** ([ch13_rwmutex_read_write_locks](./go-beyond-example-code/ch13_rwmutex_read_write_locks/)) - Line 799
- **Once: One-Time Initialization** ([ch13_once_one_time_initialization](./go-beyond-example-code/ch13_once_one_time_initialization/)) - Line 882
- **Cond: Condition Variables** ([ch13_cond_condition_variables](./go-beyond-example-code/ch13_cond_condition_variables/)) - Line 928
- **Atomic Operations** ([ch13_atomic_operations](./go-beyond-example-code/ch13_atomic_operations/)) - Line 1006
- **Examples** ([ch13_examples](./go-beyond-example-code/ch13_examples/)) - Line 1061
- **example.com",** ([ch13_examplecom](./go-beyond-example-code/ch13_examplecom/)) - Line 1136
- **Rate Limiter** ([ch13_rate_limiter](./go-beyond-example-code/ch13_rate_limiter/)) - Line 1201
- **Circuit Breaker Pattern** ([ch13_circuit_breaker_pattern](./go-beyond-example-code/ch13_circuit_breaker_pattern/)) - Line 1286
- **Worker Pool with Job Queue** ([ch13_worker_pool_with_job_queue](./go-beyond-example-code/ch13_worker_pool_with_job_queue/)) - Line 1388
- **Exercise 1** ([ch13_exercise_1](./go-beyond-example-code/ch13_exercise_1/)) - Line 1490
- **Exercise 2** ([ch13_exercise_2](./go-beyond-example-code/ch13_exercise_2/)) - Line 1557

### chapter14.md

- **Basic `select` Syntax** ([ch14_basic_select_syntax](./go-beyond-example-code/ch14_basic_select_syntax/)) - Line 38
- **Example with Multiple Selects** ([ch14_example_with_multiple_selects](./go-beyond-example-code/ch14_example_with_multiple_selects/)) - Line 84
- **Basic Multi-Channel `select`** ([ch14_basic_multi_channel_select](./go-beyond-example-code/ch14_basic_multi_channel_select/)) - Line 152
- **example shows how `select` handles multiple channels in a loop. Each iteration waits for any of the three channels to be ready, processes the message, and continues to the next iteration.** ([ch14_example_shows_how_select_handles_multiple_channels](./go-beyond-example-code/ch14_example_shows_how_select_handles_multiple_channels/)) - Line 199
- **Dynamic Channel Selection** ([ch14_dynamic_channel_selection](./go-beyond-example-code/ch14_dynamic_channel_selection/)) - Line 252
- **example shows how to dynamically process messages from multiple channels using nested `select` statements. The outer `select` handles shutdown, while the inner loop tries each channel in order.** ([ch14_example_shows_how_to_dynamically_process_messages_](./go-beyond-example-code/ch14_example_shows_how_to_dynamically_process_messages_/)) - Line 335
- **example demonstrates bidirectional communication using `select`. The goroutine continuously listens for messages on both channels, while the main function sends messages to both channels.** ([ch14_example_demonstrates_bidirectional_communication_u](./go-beyond-example-code/ch14_example_demonstrates_bidirectional_communication_u/)) - Line 371
- **Channel State Monitoring** ([ch14_channel_state_monitoring](./go-beyond-example-code/ch14_channel_state_monitoring/)) - Line 419
- **example shows how to monitor multiple channels and check their state. The `select` with `default` allows non-blocking checks of each channel's state.** ([ch14_example_shows_how_to_monitor_multiple_channels_and](./go-beyond-example-code/ch14_example_shows_how_to_monitor_multiple_channels_and/)) - Line 491
- **Advanced Non-blocking Patterns** ([ch14_advanced_non_blocking_patterns](./go-beyond-example-code/ch14_advanced_non_blocking_patterns/)) - Line 521
- **example shows how to create helper functions for non-blocking operations. The functions return boolean values to indicate success/failure, making them easy to use in conditional logic.** ([ch14_example_shows_how_to_create_helper_functions_for_n](./go-beyond-example-code/ch14_example_shows_how_to_create_helper_functions_for_n/)) - Line 577
- **example demonstrates how to process multiple channels non-blocking. The function checks each channel in a loop, processing any available messages, and only waits when all channels are empty.** ([ch14_example_demonstrates_how_to_process_multiple_chann](./go-beyond-example-code/ch14_example_demonstrates_how_to_process_multiple_chann/)) - Line 646
- **example demonstrates non-blocking operations on a buffered channel. The channel has a capacity of 2, so the first two sends succeed, but the third send fails because the channel is full.** ([ch14_example_demonstrates_non_blocking_operations_on_a_](./go-beyond-example-code/ch14_example_demonstrates_non_blocking_operations_on_a_/)) - Line 683
- **example shows how channel capacity affects non-blocking operations. Unbuffered channels (capacity 0) block on send until someone receives, while buffered channels can hold a limited number of messages.** ([ch14_example_shows_how_channel_capacity_affects_non_blo](./go-beyond-example-code/ch14_example_shows_how_channel_capacity_affects_non_blo/)) - Line 730
- **example combines non-blocking operations with timeouts. The function tries to receive from a channel but gives up after a specified timeout period.** ([ch14_example_combines_non_blocking_operations_with_time](./go-beyond-example-code/ch14_example_combines_non_blocking_operations_with_time/)) - Line 787
- **example demonstrates the basic timeout pattern using `time.After()`. The `select` statement waits for either a message from the channel or a timeout signal. Since the goroutine takes 2 seconds but the timeout is 1 second, the timeout case will execute.** ([ch14_example_demonstrates_the_basic_timeout_pattern_usi](./go-beyond-example-code/ch14_example_demonstrates_the_basic_timeout_pattern_usi/)) - Line 818
- **example shows a more sophisticated timeout pattern that handles both successful operations and errors. The function wraps any operation with a timeout and returns appropriate results or errors.** ([ch14_example_shows_a_more_sophisticated_timeout_pattern](./go-beyond-example-code/ch14_example_shows_a_more_sophisticated_timeout_pattern/)) - Line 886
- **example shows how to handle multiple operations with timeouts. The first `select` will receive from `ch1` (fast operation) and timeout on `ch2` (slow operation). The second `select` waits for the remaining slow operation.** ([ch14_example_shows_how_to_handle_multiple_operations_wi](./go-beyond-example-code/ch14_example_shows_how_to_handle_multiple_operations_wi/)) - Line 933
- **example shows how to use timeouts in a loop to process messages with a periodic timeout. The loop continues until either a message is received or a timeout occurs.** ([ch14_example_shows_how_to_use_timeouts_in_a_loop_to_pro](./go-beyond-example-code/ch14_example_shows_how_to_use_timeouts_in_a_loop_to_pro/)) - Line 969
- **example demonstrates an adaptive timeout pattern where the timeout increases when operations fail and resets when they succeed. This is useful for handling variable network conditions or processing times.** ([ch14_example_demonstrates_an_adaptive_timeout_pattern_w](./go-beyond-example-code/ch14_example_demonstrates_an_adaptive_timeout_pattern_w/)) - Line 1015
- **example shows how to combine timeouts with retry logic. The function attempts the operation multiple times with timeouts, using exponential backoff between retries.** ([ch14_example_shows_how_to_combine_timeouts_with_retry_l](./go-beyond-example-code/ch14_example_shows_how_to_combine_timeouts_with_retry_l/)) - Line 1086
- **example shows a worker pool pattern using `select`. Each worker listens for jobs on a shared channel and quit signals on a private channel. The `select` statement allows workers to handle both job processing and graceful shutdown.** ([ch14_example_shows_a_worker_pool_pattern_using_select_e](./go-beyond-example-code/ch14_example_shows_a_worker_pool_pattern_using_select_e/)) - Line 1168
- **Real-time Data Processing Pipeline** ([ch14_real_time_data_processing_pipeline](./go-beyond-example-code/ch14_real_time_data_processing_pipeline/)) - Line 1279
- **example shows a real-time data processing pipeline using `select`. Data flows from generator → processor → consumer, with each component using `select` to handle data processing and shutdown signals.** ([ch14_example_shows_a_real_time_data_processing_pipeline](./go-beyond-example-code/ch14_example_shows_a_real_time_data_processing_pipeline/)) - Line 1391
- **example demonstrates the fan-out pattern where one input channel is distributed to multiple output channels. The `select` with `default` ensures that if an output channel is full, the fan-out doesn't block and skips that channel.** ([ch14_example_demonstrates_the_fan_out_pattern_where_one](./go-beyond-example-code/ch14_example_demonstrates_the_fan_out_pattern_where_one/)) - Line 1462
- **example shows a circuit breaker pattern using `select` for timeout handling. The circuit breaker prevents cascading failures by opening the circuit when too many failures occur.** ([ch14_example_shows_a_circuit_breaker_pattern_using_sele](./go-beyond-example-code/ch14_example_shows_a_circuit_breaker_pattern_using_sele/)) - Line 1553
- **example shows a rate limiter implementation using `select`. The rate limiter uses a token bucket algorithm with `select` to handle token acquisition and refill timing.** ([ch14_example_shows_a_rate_limiter_implementation_using_](./go-beyond-example-code/ch14_example_shows_a_rate_limiter_implementation_using_/)) - Line 1650
- **Exercise 2** ([ch14_exercise_2](./go-beyond-example-code/ch14_exercise_2/)) - Line 1719
- **Exercise 3** ([ch14_exercise_3](./go-beyond-example-code/ch14_exercise_3/)) - Line 1803
- **Exercise 4** ([ch14_exercise_4](./go-beyond-example-code/ch14_exercise_4/)) - Line 1879
- **Exercise 5** ([ch14_exercise_5](./go-beyond-example-code/ch14_exercise_5/)) - Line 1979

### chapter15.md

- **Example** ([ch15_example](./go-beyond-example-code/ch15_example/)) - Line 59
- **Understanding Type vs Value vs Kind** ([ch15_understanding_type_vs_value_vs_kind](./go-beyond-example-code/ch15_understanding_type_vs_value_vs_kind/)) - Line 112
- **Working with Interface{} and Reflection** ([ch15_working_with_interface_and_reflection](./go-beyond-example-code/ch15_working_with_interface_and_reflection/)) - Line 150
- **Basic Type Operations** ([ch15_basic_type_operations](./go-beyond-example-code/ch15_basic_type_operations/)) - Line 243
- **Advanced Type Inspection** ([ch15_advanced_type_inspection](./go-beyond-example-code/ch15_advanced_type_inspection/)) - Line 360
- **Type Comparison and Conversion** ([ch15_type_comparison_and_conversion](./go-beyond-example-code/ch15_type_comparison_and_conversion/)) - Line 437
- **example** ([ch15_example_98b7e5](./go-beyond-example-code/ch15_example_98b7e5/)) - Line 514
- **Reading Values** ([ch15_reading_values](./go-beyond-example-code/ch15_reading_values/)) - Line 742
- **Advanced Value Reading** ([ch15_advanced_value_reading](./go-beyond-example-code/ch15_advanced_value_reading/)) - Line 843
- **Modifying Values** ([ch15_modifying_values](./go-beyond-example-code/ch15_modifying_values/)) - Line 947
- **example that handles type conversion and validation** ([ch15_example_that_handles_type_conversion_and_validatio](./go-beyond-example-code/ch15_example_that_handles_type_conversion_and_validatio/)) - Line 1113
- **Safe Value Modification** ([ch15_safe_value_modification](./go-beyond-example-code/ch15_safe_value_modification/)) - Line 1219
- **Basic Struct Inspection** ([ch15_basic_struct_inspection](./go-beyond-example-code/ch15_basic_struct_inspection/)) - Line 1292
- **example that shows different ways to access struct fields** ([ch15_example_that_shows_different_ways_to_access_struct](./go-beyond-example-code/ch15_example_that_shows_different_ways_to_access_struct/)) - Line 1410
- **Working with Struct Tags** ([ch15_working_with_struct_tags](./go-beyond-example-code/ch15_working_with_struct_tags/)) - Line 1614
- **Field Value Access and Modification** ([ch15_field_value_access_and_modification](./go-beyond-example-code/ch15_field_value_access_and_modification/)) - Line 1762
- **example.com}** ([ch15_examplecom](./go-beyond-example-code/ch15_examplecom/)) - Line 1877
- **Modifying Struct Fields** ([ch15_modifying_struct_fields](./go-beyond-example-code/ch15_modifying_struct_fields/)) - Line 2053
- **example.com")** ([ch15_examplecom_7346dc](./go-beyond-example-code/ch15_examplecom_7346dc/)) - Line 2106
- **example.com")** ([ch15_examplecom_b66c7d](./go-beyond-example-code/ch15_examplecom_b66c7d/)) - Line 2210
- **Examples** ([ch15_examples](./go-beyond-example-code/ch15_examples/)) - Line 2327
- **example.com","active"** ([ch15_examplecomactive](./go-beyond-example-code/ch15_examplecomactive/)) - Line 2572
- **Configuration Loader** ([ch15_configuration_loader](./go-beyond-example-code/ch15_configuration_loader/)) - Line 2817
- **Exercise 1** ([ch15_exercise_1](./go-beyond-example-code/ch15_exercise_1/)) - Line 2908
- **Exercise 2** ([ch15_exercise_2](./go-beyond-example-code/ch15_exercise_2/)) - Line 3061
- **Exercise 3** ([ch15_exercise_3](./go-beyond-example-code/ch15_exercise_3/)) - Line 3295

### chapter16.md

- **Example() {** ([ch16_example](./go-beyond-example-code/ch16_example/)) - Line 66
- **Basic Mutex Usage** ([ch16_basic_mutex_usage](./go-beyond-example-code/ch16_basic_mutex_usage/)) - Line 119
- **Mutex Best Practices** ([ch16_mutex_best_practices](./go-beyond-example-code/ch16_mutex_best_practices/)) - Line 176
- **Basic RWMutex Usage** ([ch16_basic_rwmutex_usage](./go-beyond-example-code/ch16_basic_rwmutex_usage/)) - Line 436
- **RWMutex vs Mutex Performance** ([ch16_rwmutex_vs_mutex_performance](./go-beyond-example-code/ch16_rwmutex_vs_mutex_performance/)) - Line 525
- **Basic WaitGroup Usage** ([ch16_basic_waitgroup_usage](./go-beyond-example-code/ch16_basic_waitgroup_usage/)) - Line 829
- **WaitGroup with Results** ([ch16_waitgroup_with_results](./go-beyond-example-code/ch16_waitgroup_with_results/)) - Line 863
- **1. WaitGroup with Error Handling** ([ch16_1_waitgroup_with_error_handling](./go-beyond-example-code/ch16_1_waitgroup_with_error_handling/)) - Line 924
- **2. WaitGroup with Timeout** ([ch16_2_waitgroup_with_timeout](./go-beyond-example-code/ch16_2_waitgroup_with_timeout/)) - Line 993
- **3. WaitGroup with Dynamic Workers** ([ch16_3_waitgroup_with_dynamic_workers](./go-beyond-example-code/ch16_3_waitgroup_with_dynamic_workers/)) - Line 1045
- **Basic Once Usage** ([ch16_basic_once_usage](./go-beyond-example-code/ch16_basic_once_usage/)) - Line 1199
- **Once with Error Handling** ([ch16_once_with_error_handling](./go-beyond-example-code/ch16_once_with_error_handling/)) - Line 1249
- **1. Once with Error Handling and Retry** ([ch16_1_once_with_error_handling_and_retry](./go-beyond-example-code/ch16_1_once_with_error_handling_and_retry/)) - Line 1310
- **2. Once with Configuration** ([ch16_2_once_with_configuration](./go-beyond-example-code/ch16_2_once_with_configuration/)) - Line 1376
- **3. Once with Resource Cleanup** ([ch16_3_once_with_resource_cleanup](./go-beyond-example-code/ch16_3_once_with_resource_cleanup/)) - Line 1432
- **Basic Cond Usage** ([ch16_basic_cond_usage](./go-beyond-example-code/ch16_basic_cond_usage/)) - Line 1586
- **Cond with Broadcast** ([ch16_cond_with_broadcast](./go-beyond-example-code/ch16_cond_with_broadcast/)) - Line 1664
- **1. Worker Pool with Job Queue** ([ch16_1_worker_pool_with_job_queue](./go-beyond-example-code/ch16_1_worker_pool_with_job_queue/)) - Line 1733
- **2. Resource Pool with Condition Variables** ([ch16_2_resource_pool_with_condition_variables](./go-beyond-example-code/ch16_2_resource_pool_with_condition_variables/)) - Line 1839
- **3. Barrier Synchronization** ([ch16_3_barrier_synchronization](./go-beyond-example-code/ch16_3_barrier_synchronization/)) - Line 1923
- **Basic Atomic Operations** ([ch16_basic_atomic_operations](./go-beyond-example-code/ch16_basic_atomic_operations/)) - Line 2078
- **Atomic vs Mutex Performance** ([ch16_atomic_vs_mutex_performance](./go-beyond-example-code/ch16_atomic_vs_mutex_performance/)) - Line 2126
- **1. Atomic Flags and State Management** ([ch16_1_atomic_flags_and_state_management](./go-beyond-example-code/ch16_1_atomic_flags_and_state_management/)) - Line 2216
- **2. Atomic Reference Counting** ([ch16_2_atomic_reference_counting](./go-beyond-example-code/ch16_2_atomic_reference_counting/)) - Line 2293
- **3. Atomic Compare and Swap (CAS)** ([ch16_3_atomic_compare_and_swap_cas](./go-beyond-example-code/ch16_3_atomic_compare_and_swap_cas/)) - Line 2360
- **Exercise 1** ([ch16_exercise_1](./go-beyond-example-code/ch16_exercise_1/)) - Line 2544
- **Exercise 2** ([ch16_exercise_2](./go-beyond-example-code/ch16_exercise_2/)) - Line 2632
- **Workgroups** ([ch16_workgroups](./go-beyond-example-code/ch16_workgroups/)) - Line 2713
- **Workgroup with Results** ([ch16_workgroup_with_results](./go-beyond-example-code/ch16_workgroup_with_results/)) - Line 2770
- **Errorgroups** ([ch16_errorgroups](./go-beyond-example-code/ch16_errorgroups/)) - Line 2849
- **Combined Workgroup and Errorgroup** ([ch16_combined_workgroup_and_errorgroup](./go-beyond-example-code/ch16_combined_workgroup_and_errorgroup/)) - Line 2934

### chapter17.md

- **Basic Context Usage** ([ch17_basic_context_usage](./go-beyond-example-code/ch17_basic_context_usage/)) - Line 51
- **Basic Timeout** ([ch17_basic_timeout](./go-beyond-example-code/ch17_basic_timeout/)) - Line 87
- **Timeout with Different Scenarios** ([ch17_timeout_with_different_scenarios](./go-beyond-example-code/ch17_timeout_with_different_scenarios/)) - Line 132
- **Multiple Timeouts** ([ch17_multiple_timeouts](./go-beyond-example-code/ch17_multiple_timeouts/)) - Line 175
- **Nested Timeouts** ([ch17_nested_timeouts](./go-beyond-example-code/ch17_nested_timeouts/)) - Line 216
- **Manual Cancellation** ([ch17_manual_cancellation](./go-beyond-example-code/ch17_manual_cancellation/)) - Line 280
- **Cancellation with User Input** ([ch17_cancellation_with_user_input](./go-beyond-example-code/ch17_cancellation_with_user_input/)) - Line 331
- **Conditional Cancellation** ([ch17_conditional_cancellation](./go-beyond-example-code/ch17_conditional_cancellation/)) - Line 381
- **Cancellation with Deadline** ([ch17_cancellation_with_deadline](./go-beyond-example-code/ch17_cancellation_with_deadline/)) - Line 434
- **Cancellation Chain** ([ch17_cancellation_chain](./go-beyond-example-code/ch17_cancellation_chain/)) - Line 480
- **Basic Value Storage** ([ch17_basic_value_storage](./go-beyond-example-code/ch17_basic_value_storage/)) - Line 548
- **Value Propagation Through Functions** ([ch17_value_propagation_through_functions](./go-beyond-example-code/ch17_value_propagation_through_functions/)) - Line 592
- **Context Value Inheritance** ([ch17_context_value_inheritance](./go-beyond-example-code/ch17_context_value_inheritance/)) - Line 635
- **Type-Safe Context Values** ([ch17_type_safe_context_values](./go-beyond-example-code/ch17_type_safe_context_values/)) - Line 688
- **Advanced Context Value Patterns** ([ch17_advanced_context_value_patterns](./go-beyond-example-code/ch17_advanced_context_value_patterns/)) - Line 748
- **Context Value Validation** ([ch17_context_value_validation](./go-beyond-example-code/ch17_context_value_validation/)) - Line 808
- **Passing Context Through Functions** ([ch17_passing_context_through_functions](./go-beyond-example-code/ch17_passing_context_through_functions/)) - Line 895
- **Multi-Layer Context Propagation** ([ch17_multi_layer_context_propagation](./go-beyond-example-code/ch17_multi_layer_context_propagation/)) - Line 946
- **Context Propagation with Goroutines** ([ch17_context_propagation_with_goroutines](./go-beyond-example-code/ch17_context_propagation_with_goroutines/)) - Line 1005
- **Context Propagation with Error Handling** ([ch17_context_propagation_with_error_handling](./go-beyond-example-code/ch17_context_propagation_with_error_handling/)) - Line 1074
- **Context in HTTP Handlers** ([ch17_context_in_http_handlers](./go-beyond-example-code/ch17_context_in_http_handlers/)) - Line 1166
- **Examples** ([ch17_examples](./go-beyond-example-code/ch17_examples/)) - Line 1235
- **Advanced Database Operations with Context** ([ch17_advanced_database_operations_with_context](./go-beyond-example-code/ch17_advanced_database_operations_with_context/)) - Line 1294
- **Microservice Communication with Context** ([ch17_microservice_communication_with_context](./go-beyond-example-code/ch17_microservice_communication_with_context/)) - Line 1378
- **API Client with Context** ([ch17_api_client_with_context](./go-beyond-example-code/ch17_api_client_with_context/)) - Line 1454
- **Worker Pool with Context** ([ch17_worker_pool_with_context](./go-beyond-example-code/ch17_worker_pool_with_context/)) - Line 1513
- **Advanced Worker Pool with Context** ([ch17_advanced_worker_pool_with_context](./go-beyond-example-code/ch17_advanced_worker_pool_with_context/)) - Line 1600
- **Context-Aware Rate Limiter** ([ch17_context_aware_rate_limiter](./go-beyond-example-code/ch17_context_aware_rate_limiter/)) - Line 1727
- **Exercise 1** ([ch17_exercise_1](./go-beyond-example-code/ch17_exercise_1/)) - Line 1808
- **Exercise 2** ([ch17_exercise_2](./go-beyond-example-code/ch17_exercise_2/)) - Line 1859

### chapter19.md

- **Basic Math Operations** ([ch19_basic_math_operations](./go-beyond-example-code/ch19_basic_math_operations/)) - Line 33
- **Advanced Math Operations** ([ch19_advanced_math_operations](./go-beyond-example-code/ch19_advanced_math_operations/)) - Line 86
- **Mathematical Constants** ([ch19_mathematical_constants](./go-beyond-example-code/ch19_mathematical_constants/)) - Line 123
- **Basic Trigonometry** ([ch19_basic_trigonometry](./go-beyond-example-code/ch19_basic_trigonometry/)) - Line 163
- **Advanced Trigonometric Applications** ([ch19_advanced_trigonometric_applications](./go-beyond-example-code/ch19_advanced_trigonometric_applications/)) - Line 230
- **Example** ([ch19_example](./go-beyond-example-code/ch19_example/)) - Line 360
- **Basic Log and Exp Operations** ([ch19_basic_log_and_exp_operations](./go-beyond-example-code/ch19_basic_log_and_exp_operations/)) - Line 421
- **Advanced Logarithmic and Exponential Applications** ([ch19_advanced_logarithmic_and_exponential_applications](./go-beyond-example-code/ch19_advanced_logarithmic_and_exponential_applications/)) - Line 495
- **Examples** ([ch19_examples](./go-beyond-example-code/ch19_examples/)) - Line 667
- **Basic Random Numbers** ([ch19_basic_random_numbers](./go-beyond-example-code/ch19_basic_random_numbers/)) - Line 731
- **Advanced Random Number Generation** ([ch19_advanced_random_number_generation](./go-beyond-example-code/ch19_advanced_random_number_generation/)) - Line 792
- **Advanced Random Number Generation** ([ch19_advanced_random_number_generation_c1cbff](./go-beyond-example-code/ch19_advanced_random_number_generation_c1cbff/)) - Line 987
- **Prime Number Algorithms** ([ch19_prime_number_algorithms](./go-beyond-example-code/ch19_prime_number_algorithms/)) - Line 1068
- **Fibonacci Sequence** ([ch19_fibonacci_sequence](./go-beyond-example-code/ch19_fibonacci_sequence/)) - Line 1302
- **Greatest Common Divisor** ([ch19_greatest_common_divisor](./go-beyond-example-code/ch19_greatest_common_divisor/)) - Line 1540
- **Exercise 1** ([ch19_exercise_1](./go-beyond-example-code/ch19_exercise_1/)) - Line 1595
- **Exercise 2** ([ch19_exercise_2](./go-beyond-example-code/ch19_exercise_2/)) - Line 1878

### chapter20.md

- **Method 1: `ioutil.ReadFile` (Simple but Memory-Intensive)** ([ch20_method_1_ioutilreadfile_simple_but_memory_intensiv](./go-beyond-example-code/ch20_method_1_ioutilreadfile_simple_but_memory_intensiv/)) - Line 31
- **example.txt")** ([ch20_exampletxt](./go-beyond-example-code/ch20_exampletxt/)) - Line 60
- **Method 3: `os.ReadFile` (Go 1.16+ Recommended)** ([ch20_method_3_osreadfile_go_116_recommended](./go-beyond-example-code/ch20_method_3_osreadfile_go_116_recommended/)) - Line 104
- **example.txt")** ([ch20_exampletxt_ce9959](./go-beyond-example-code/ch20_exampletxt_ce9959/)) - Line 130
- **example.txt")** ([ch20_exampletxt_c81b28](./go-beyond-example-code/ch20_exampletxt_c81b28/)) - Line 159
- **example.txt", 1024*1024) // 1MB limit** ([ch20_exampletxt_10241024_1mb_limit](./go-beyond-example-code/ch20_exampletxt_10241024_1mb_limit/)) - Line 199
- **Basic Line-by-Line Reading** ([ch20_basic_line_by_line_reading](./go-beyond-example-code/ch20_basic_line_by_line_reading/)) - Line 235
- **Advanced Scanner Configuration** ([ch20_advanced_scanner_configuration](./go-beyond-example-code/ch20_advanced_scanner_configuration/)) - Line 274
- **Word-by-Word Reading** ([ch20_word_by_word_reading](./go-beyond-example-code/ch20_word_by_word_reading/)) - Line 321
- **Custom Delimiter Functions** ([ch20_custom_delimiter_functions](./go-beyond-example-code/ch20_custom_delimiter_functions/)) - Line 368
- **Processing Lines with Context** ([ch20_processing_lines_with_context](./go-beyond-example-code/ch20_processing_lines_with_context/)) - Line 423
- **Reading with Line Filtering** ([ch20_reading_with_line_filtering](./go-beyond-example-code/ch20_reading_with_line_filtering/)) - Line 485
- **example.txt", containsFilter("error"))** ([ch20_exampletxt_containsfiltererror](./go-beyond-example-code/ch20_exampletxt_containsfiltererror/)) - Line 558
- **Basic Path Manipulation** ([ch20_basic_path_manipulation](./go-beyond-example-code/ch20_basic_path_manipulation/)) - Line 625
- **Advanced Path Operations** ([ch20_advanced_path_operations](./go-beyond-example-code/ch20_advanced_path_operations/)) - Line 673
- **Path Validation and Sanitization** ([ch20_path_validation_and_sanitization](./go-beyond-example-code/ch20_path_validation_and_sanitization/)) - Line 720
- **Listing Directory Contents** ([ch20_listing_directory_contents](./go-beyond-example-code/ch20_listing_directory_contents/)) - Line 795
- **Recursive Directory Walking** ([ch20_recursive_directory_walking](./go-beyond-example-code/ch20_recursive_directory_walking/)) - Line 871
- **Filtered Directory Walking** ([ch20_filtered_directory_walking](./go-beyond-example-code/ch20_filtered_directory_walking/)) - Line 951
- **Finding Files by Pattern** ([ch20_finding_files_by_pattern](./go-beyond-example-code/ch20_finding_files_by_pattern/)) - Line 1050
- **Basic CSV Reading** ([ch20_basic_csv_reading](./go-beyond-example-code/ch20_basic_csv_reading/)) - Line 1146
- **Advanced CSV Processing** ([ch20_advanced_csv_processing](./go-beyond-example-code/ch20_advanced_csv_processing/)) - Line 1179
- **Streaming CSV Processing** ([ch20_streaming_csv_processing](./go-beyond-example-code/ch20_streaming_csv_processing/)) - Line 1289
- **Basic JSON Reading** ([ch20_basic_json_reading](./go-beyond-example-code/ch20_basic_json_reading/)) - Line 1342
- **Advanced JSON Processing** ([ch20_advanced_json_processing](./go-beyond-example-code/ch20_advanced_json_processing/)) - Line 1384
- **Streaming JSON Processing** ([ch20_streaming_json_processing](./go-beyond-example-code/ch20_streaming_json_processing/)) - Line 1469
- **Word Frequency Analysis** ([ch20_word_frequency_analysis](./go-beyond-example-code/ch20_word_frequency_analysis/)) - Line 1531
- **Log File Processing** ([ch20_log_file_processing](./go-beyond-example-code/ch20_log_file_processing/)) - Line 1652
- **Examples** ([ch20_examples](./go-beyond-example-code/ch20_examples/)) - Line 1791
- **Comprehensive Log File Analyzer** ([ch20_comprehensive_log_file_analyzer](./go-beyond-example-code/ch20_comprehensive_log_file_analyzer/)) - Line 2020
- **File Backup Utility** ([ch20_file_backup_utility](./go-beyond-example-code/ch20_file_backup_utility/)) - Line 2339
- **Exercise 1** ([ch20_exercise_1](./go-beyond-example-code/ch20_exercise_1/)) - Line 2662
- **Exercise 2** ([ch20_exercise_2](./go-beyond-example-code/ch20_exercise_2/)) - Line 2714
- **example.txt")** ([ch20_exampletxt_d033b7](./go-beyond-example-code/ch20_exampletxt_d033b7/)) - Line 2785
- **Custom Error Types** ([ch20_custom_error_types](./go-beyond-example-code/ch20_custom_error_types/)) - Line 2828
- **example.txt")** ([ch20_exampletxt_6216ce](./go-beyond-example-code/ch20_exampletxt_6216ce/)) - Line 2901
- **Memory-Efficient Processing** ([ch20_memory_efficient_processing](./go-beyond-example-code/ch20_memory_efficient_processing/)) - Line 2963
- **Path Validation** ([ch20_path_validation](./go-beyond-example-code/ch20_path_validation/)) - Line 3025
- **File Size Limits** ([ch20_file_size_limits](./go-beyond-example-code/ch20_file_size_limits/)) - Line 3086
- **example.txt", 1024*1024) // 1MB limit** ([ch20_exampletxt_10241024_1mb_limit_d3b000](./go-beyond-example-code/ch20_exampletxt_10241024_1mb_limit_d3b000/)) - Line 3128
- **Mock File System for Testing** ([ch20_mock_file_system_for_testing](./go-beyond-example-code/ch20_mock_file_system_for_testing/)) - Line 3197

### chapter21.md

- **Example** ([ch21_example](./go-beyond-example-code/ch21_example/)) - Line 53
- **Basic Actions** ([ch21_basic_actions](./go-beyond-example-code/ch21_basic_actions/)) - Line 119
- **Template Variables** ([ch21_template_variables](./go-beyond-example-code/ch21_template_variables/)) - Line 206
- **Advanced Variable Usage** ([ch21_advanced_variable_usage](./go-beyond-example-code/ch21_advanced_variable_usage/)) - Line 280
- **Example |** ([ch21_example_07d479](./go-beyond-example-code/ch21_example_07d479/)) - Line 376
- **example.com",** ([ch21_examplecom](./go-beyond-example-code/ch21_examplecom/)) - Line 459
- **Custom Functions** ([ch21_custom_functions](./go-beyond-example-code/ch21_custom_functions/)) - Line 555
- **example.com",** ([ch21_examplecom_00dbab](./go-beyond-example-code/ch21_examplecom_00dbab/)) - Line 645
- **Basic HTML Template** ([ch21_basic_html_template](./go-beyond-example-code/ch21_basic_html_template/)) - Line 798

### chapter22.md

- **Basic Generic Function** ([ch22_basic_generic_function](./go-beyond-example-code/ch22_basic_generic_function/)) - Line 73
- **Basic Generic Functions** ([ch22_basic_generic_functions](./go-beyond-example-code/ch22_basic_generic_functions/)) - Line 120
- **Generic Functions with Constraints** ([ch22_generic_functions_with_constraints](./go-beyond-example-code/ch22_generic_functions_with_constraints/)) - Line 258
- **Generic Structs** ([ch22_generic_structs](./go-beyond-example-code/ch22_generic_structs/)) - Line 445
- **Generic Queue Implementation** ([ch22_generic_queue_implementation](./go-beyond-example-code/ch22_generic_queue_implementation/)) - Line 584
- **Generic Maps** ([ch22_generic_maps](./go-beyond-example-code/ch22_generic_maps/)) - Line 690
- **Generic Linked List** ([ch22_generic_linked_list](./go-beyond-example-code/ch22_generic_linked_list/)) - Line 846
- **Basic Type Constraints** ([ch22_basic_type_constraints](./go-beyond-example-code/ch22_basic_type_constraints/)) - Line 997
- **Advanced Type Constraints** ([ch22_advanced_type_constraints](./go-beyond-example-code/ch22_advanced_type_constraints/)) - Line 1165
- **Exercise 1** ([ch22_exercise_1](./go-beyond-example-code/ch22_exercise_1/)) - Line 1402
- **main_program** ([ch22_main_program](./go-beyond-example-code/ch22_main_program/)) - Line 1616

### chapter23.md

- **Basic Generic Slice Functions** ([ch23_basic_generic_slice_functions](./go-beyond-example-code/ch23_basic_generic_slice_functions/)) - Line 21
- **Advanced Generic Slice Utilities** ([ch23_advanced_generic_slice_utilities](./go-beyond-example-code/ch23_advanced_generic_slice_utilities/)) - Line 147
- **Generic Slice Transformations** ([ch23_generic_slice_transformations](./go-beyond-example-code/ch23_generic_slice_transformations/)) - Line 260
- **Generic Stack** ([ch23_generic_stack](./go-beyond-example-code/ch23_generic_stack/)) - Line 347
- **Generic Queue** ([ch23_generic_queue](./go-beyond-example-code/ch23_generic_queue/)) - Line 475
- **Generic Priority Queue** ([ch23_generic_priority_queue](./go-beyond-example-code/ch23_generic_priority_queue/)) - Line 605
- **Generic Sorting** ([ch23_generic_sorting](./go-beyond-example-code/ch23_generic_sorting/)) - Line 726
- **Generic Search Algorithms** ([ch23_generic_search_algorithms](./go-beyond-example-code/ch23_generic_search_algorithms/)) - Line 895
- **Generic Set Operations** ([ch23_generic_set_operations](./go-beyond-example-code/ch23_generic_set_operations/)) - Line 989
- **Exercise 1** ([ch23_exercise_1](./go-beyond-example-code/ch23_exercise_1/)) - Line 1133
- **Exercise 2** ([ch23_exercise_2](./go-beyond-example-code/ch23_exercise_2/)) - Line 1229
- **Exercise 3** ([ch23_exercise_3](./go-beyond-example-code/ch23_exercise_3/)) - Line 1375
- **Exercise 2** ([ch23_exercise_2_017828](./go-beyond-example-code/ch23_exercise_2_017828/)) - Line 1508

### chapter24.md

- **Multiple Embedding and Method Resolution** ([ch24_multiple_embedding_and_method_resolution](./go-beyond-example-code/ch24_multiple_embedding_and_method_resolution/)) - Line 31
- **Method Resolution and Conflicts** ([ch24_method_resolution_and_conflicts](./go-beyond-example-code/ch24_method_resolution_and_conflicts/)) - Line 140
- **Embedding with Pointer Types** ([ch24_embedding_with_pointer_types](./go-beyond-example-code/ch24_embedding_with_pointer_types/)) - Line 201
- **Interface Embedding with Constraints** ([ch24_interface_embedding_with_constraints](./go-beyond-example-code/ch24_interface_embedding_with_constraints/)) - Line 261
- **Generic Workgroup with Type Safety** ([ch24_generic_workgroup_with_type_safety](./go-beyond-example-code/ch24_generic_workgroup_with_type_safety/)) - Line 349
- **Advanced Errorgroup with Error Aggregation** ([ch24_advanced_errorgroup_with_error_aggregation](./go-beyond-example-code/ch24_advanced_errorgroup_with_error_aggregation/)) - Line 470
- **Workgroup with Progress Tracking** ([ch24_workgroup_with_progress_tracking](./go-beyond-example-code/ch24_workgroup_with_progress_tracking/)) - Line 644
- **Pipeline with Workgroups** ([ch24_pipeline_with_workgroups](./go-beyond-example-code/ch24_pipeline_with_workgroups/)) - Line 786
- **Dynamic Method Invocation** ([ch24_dynamic_method_invocation](./go-beyond-example-code/ch24_dynamic_method_invocation/)) - Line 927
- **Advanced Struct Field Manipulation** ([ch24_advanced_struct_field_manipulation](./go-beyond-example-code/ch24_advanced_struct_field_manipulation/)) - Line 1060
- **Dynamic Struct Creation and Manipulation** ([ch24_dynamic_struct_creation_and_manipulation](./go-beyond-example-code/ch24_dynamic_struct_creation_and_manipulation/)) - Line 1265
- **Advanced Type Introspection** ([ch24_advanced_type_introspection](./go-beyond-example-code/ch24_advanced_type_introspection/)) - Line 1455
- **Advanced Object Pool Pattern** ([ch24_advanced_object_pool_pattern](./go-beyond-example-code/ch24_advanced_object_pool_pattern/)) - Line 1620
- **Advanced Lazy Loading Pattern** ([ch24_advanced_lazy_loading_pattern](./go-beyond-example-code/ch24_advanced_lazy_loading_pattern/)) - Line 1764
- **Memory-Efficient String Building** ([ch24_memory_efficient_string_building](./go-beyond-example-code/ch24_memory_efficient_string_building/)) - Line 1902
- **Lazy Loading Pattern** ([ch24_lazy_loading_pattern](./go-beyond-example-code/ch24_lazy_loading_pattern/)) - Line 2016
- **Property-Based Testing with Custom Generators** ([ch24_property_based_testing_with_custom_generators](./go-beyond-example-code/ch24_property_based_testing_with_custom_generators/)) - Line 2099
- **Advanced Mock Testing with Interfaces** ([ch24_advanced_mock_testing_with_interfaces](./go-beyond-example-code/ch24_advanced_mock_testing_with_interfaces/)) - Line 2255
- **example.com", 30)** ([ch24_examplecom_30](./go-beyond-example-code/ch24_examplecom_30/)) - Line 2508
- **Mock Testing with Interfaces** ([ch24_mock_testing_with_interfaces](./go-beyond-example-code/ch24_mock_testing_with_interfaces/)) - Line 2618
- **Advanced Circuit Breaker Pattern** ([ch24_advanced_circuit_breaker_pattern](./go-beyond-example-code/ch24_advanced_circuit_breaker_pattern/)) - Line 2746
- **Advanced Rate Limiter Pattern** ([ch24_advanced_rate_limiter_pattern](./go-beyond-example-code/ch24_advanced_rate_limiter_pattern/)) - Line 2920
- **Retry Pattern with Exponential Backoff** ([ch24_retry_pattern_with_exponential_backoff](./go-beyond-example-code/ch24_retry_pattern_with_exponential_backoff/)) - Line 3088
- **Rate Limiter Pattern** ([ch24_rate_limiter_pattern](./go-beyond-example-code/ch24_rate_limiter_pattern/)) - Line 3236

### chapter3.md

- **examples** ([ch03_examples](./go-beyond-example-code/ch03_examples/)) - Line 24

### chapter4.md

- **Examples** ([ch04_examples](./go-beyond-example-code/ch04_examples/)) - Line 55
- **Examples** ([ch04_examples_83b213](./go-beyond-example-code/ch04_examples_83b213/)) - Line 108
- **Real-World Integer Usage** ([ch04_real_world_integer_usage](./go-beyond-example-code/ch04_real_world_integer_usage/)) - Line 151
- **Examples** ([ch04_examples_ebc10e](./go-beyond-example-code/ch04_examples_ebc10e/)) - Line 228
- **Floating-Point Precision Issues** ([ch04_floating_point_precision_issues](./go-beyond-example-code/ch04_floating_point_precision_issues/)) - Line 267
- **example** ([ch04_example](./go-beyond-example-code/ch04_example/)) - Line 306
- **Special Float Values** ([ch04_special_float_values](./go-beyond-example-code/ch04_special_float_values/)) - Line 335
- **Real-World Float Usage** ([ch04_real_world_float_usage](./go-beyond-example-code/ch04_real_world_float_usage/)) - Line 366
- **String Basics** ([ch04_string_basics](./go-beyond-example-code/ch04_string_basics/)) - Line 423
- **String Internals and UTF-8** ([ch04_string_internals_and_utf_8](./go-beyond-example-code/ch04_string_internals_and_utf_8/)) - Line 447
- **Character Types** ([ch04_character_types](./go-beyond-example-code/ch04_character_types/)) - Line 482
- **Examples** ([ch04_examples_136436](./go-beyond-example-code/ch04_examples_136436/)) - Line 514
- **Unicode and Internationalization** ([ch04_unicode_and_internationalization](./go-beyond-example-code/ch04_unicode_and_internationalization/)) - Line 551
- **String Building and Performance** ([ch04_string_building_and_performance](./go-beyond-example-code/ch04_string_building_and_performance/)) - Line 592
- **Real-World String Processing** ([ch04_real_world_string_processing](./go-beyond-example-code/ch04_real_world_string_processing/)) - Line 625
- **Integer Conversions** ([ch04_integer_conversions](./go-beyond-example-code/ch04_integer_conversions/)) - Line 699
- **Dangerous Conversions and Data Loss** ([ch04_dangerous_conversions_and_data_loss](./go-beyond-example-code/ch04_dangerous_conversions_and_data_loss/)) - Line 726
- **Float Conversions** ([ch04_float_conversions](./go-beyond-example-code/ch04_float_conversions/)) - Line 758
- **String Conversions with Error Handling** ([ch04_string_conversions_with_error_handling](./go-beyond-example-code/ch04_string_conversions_with_error_handling/)) - Line 794
- **Advanced String Conversions** ([ch04_advanced_string_conversions](./go-beyond-example-code/ch04_advanced_string_conversions/)) - Line 846
- **Type Conversion Best Practices** ([ch04_type_conversion_best_practices](./go-beyond-example-code/ch04_type_conversion_best_practices/)) - Line 896
- **Examples** ([ch04_examples_3c15cb](./go-beyond-example-code/ch04_examples_3c15cb/)) - Line 961
- **Basic Arithmetic** ([ch04_basic_arithmetic](./go-beyond-example-code/ch04_basic_arithmetic/)) - Line 1050
- **Integer Division vs Float Division** ([ch04_integer_division_vs_float_division](./go-beyond-example-code/ch04_integer_division_vs_float_division/)) - Line 1073
- **Advanced Math Operations** ([ch04_advanced_math_operations](./go-beyond-example-code/ch04_advanced_math_operations/)) - Line 1118
- **Mathematical Constants and Special Values** ([ch04_mathematical_constants_and_special_values](./go-beyond-example-code/ch04_mathematical_constants_and_special_values/)) - Line 1142
- **Trigonometric Functions** ([ch04_trigonometric_functions](./go-beyond-example-code/ch04_trigonometric_functions/)) - Line 1170
- **Logarithmic and Exponential Functions** ([ch04_logarithmic_and_exponential_functions](./go-beyond-example-code/ch04_logarithmic_and_exponential_functions/)) - Line 1200
- **Edge Cases and Special Situations** ([ch04_edge_cases_and_special_situations](./go-beyond-example-code/ch04_edge_cases_and_special_situations/)) - Line 1236
- **Real-World Mathematical Applications** ([ch04_real_world_mathematical_applications](./go-beyond-example-code/ch04_real_world_mathematical_applications/)) - Line 1287
- **Performance Considerations** ([ch04_performance_considerations](./go-beyond-example-code/ch04_performance_considerations/)) - Line 1368
- **Binary, Octal, and Hexadecimal** ([ch04_binary_octal_and_hexadecimal](./go-beyond-example-code/ch04_binary_octal_and_hexadecimal/)) - Line 1447
- **Number Base Conversion** ([ch04_number_base_conversion](./go-beyond-example-code/ch04_number_base_conversion/)) - Line 1475
- **Bitwise Operations** ([ch04_bitwise_operations](./go-beyond-example-code/ch04_bitwise_operations/)) - Line 1514
- **Advanced Bitwise Operations** ([ch04_advanced_bitwise_operations](./go-beyond-example-code/ch04_advanced_bitwise_operations/)) - Line 1535
- **Practical Bitwise Applications** ([ch04_practical_bitwise_applications](./go-beyond-example-code/ch04_practical_bitwise_applications/)) - Line 1578
- **Color and Graphics Applications** ([ch04_color_and_graphics_applications](./go-beyond-example-code/ch04_color_and_graphics_applications/)) - Line 1616
- **Network and Protocol Applications** ([ch04_network_and_protocol_applications](./go-beyond-example-code/ch04_network_and_protocol_applications/)) - Line 1662
- **Bit Fields and Flags** ([ch04_bit_fields_and_flags](./go-beyond-example-code/ch04_bit_fields_and_flags/)) - Line 1708
- **Performance and Optimization** ([ch04_performance_and_optimization](./go-beyond-example-code/ch04_performance_and_optimization/)) - Line 1760
- **Examples** ([ch04_examples_2069a1](./go-beyond-example-code/ch04_examples_2069a1/)) - Line 1825
- **Number Guessing Game** ([ch04_number_guessing_game](./go-beyond-example-code/ch04_number_guessing_game/)) - Line 1883
- **Exercise 1** ([ch04_exercise_1](./go-beyond-example-code/ch04_exercise_1/)) - Line 1929
- **Exercise 1** ([ch04_exercise_1_44b768](./go-beyond-example-code/ch04_exercise_1_44b768/)) - Line 1947
- **Exercise 3** ([ch04_exercise_3](./go-beyond-example-code/ch04_exercise_3/)) - Line 1980
- **Exercise 4** ([ch04_exercise_4](./go-beyond-example-code/ch04_exercise_4/)) - Line 2009
- **Exercise 5** ([ch04_exercise_5](./go-beyond-example-code/ch04_exercise_5/)) - Line 2047
- **Exercise 6** ([ch04_exercise_6](./go-beyond-example-code/ch04_exercise_6/)) - Line 2103
- **Exercise 7** ([ch04_exercise_7](./go-beyond-example-code/ch04_exercise_7/)) - Line 2140
- **Exercise 8** ([ch04_exercise_8](./go-beyond-example-code/ch04_exercise_8/)) - Line 2204
- **Exercise 9** ([ch04_exercise_9](./go-beyond-example-code/ch04_exercise_9/)) - Line 2293
- **Exercise 10** ([ch04_exercise_10](./go-beyond-example-code/ch04_exercise_10/)) - Line 2343

### chapter5.md

- **Basic `for` Loop** ([ch05_basic_for_loop](./go-beyond-example-code/ch05_basic_for_loop/)) - Line 21
- **Example | When It Runs |** ([ch05_example_when_it_runs](./go-beyond-example-code/ch05_example_when_it_runs/)) - Line 78
- **1. Traditional C-style Loop** ([ch05_1_traditional_c_style_loop](./go-beyond-example-code/ch05_1_traditional_c_style_loop/)) - Line 132
- **2. While-style Loop** ([ch05_2_while_style_loop](./go-beyond-example-code/ch05_2_while_style_loop/)) - Line 157
- **example** ([ch05_example](./go-beyond-example-code/ch05_example/)) - Line 181
- **Range with Strings** ([ch05_range_with_strings](./go-beyond-example-code/ch05_range_with_strings/)) - Line 209
- **Understanding Range with Strings** ([ch05_understanding_range_with_strings](./go-beyond-example-code/ch05_understanding_range_with_strings/)) - Line 281
- **Range with Arrays and Slices** ([ch05_range_with_arrays_and_slices](./go-beyond-example-code/ch05_range_with_arrays_and_slices/)) - Line 335
- **Advanced Range with Slices** ([ch05_advanced_range_with_slices](./go-beyond-example-code/ch05_advanced_range_with_slices/)) - Line 418
- **Range with Maps** ([ch05_range_with_maps](./go-beyond-example-code/ch05_range_with_maps/)) - Line 490
- **Advanced Range with Maps** ([ch05_advanced_range_with_maps](./go-beyond-example-code/ch05_advanced_range_with_maps/)) - Line 563
- **`break` Statement** ([ch05_break_statement](./go-beyond-example-code/ch05_break_statement/)) - Line 667
- **Examples** ([ch05_examples](./go-beyond-example-code/ch05_examples/)) - Line 763
- **`continue` Statement** ([ch05_continue_statement](./go-beyond-example-code/ch05_continue_statement/)) - Line 841
- **Examples** ([ch05_examples_315882](./go-beyond-example-code/ch05_examples_315882/)) - Line 931
- **Break vs Continue Comparison** ([ch05_break_vs_continue_comparison](./go-beyond-example-code/ch05_break_vs_continue_comparison/)) - Line 1032
- **Basic Nested Loops** ([ch05_basic_nested_loops](./go-beyond-example-code/ch05_basic_nested_loops/)) - Line 1097
- **Advanced Nested Loop Patterns** ([ch05_advanced_nested_loop_patterns](./go-beyond-example-code/ch05_advanced_nested_loop_patterns/)) - Line 1168
- **Real-World Nested Loop Applications** ([ch05_real_world_nested_loop_applications](./go-beyond-example-code/ch05_real_world_nested_loop_applications/)) - Line 1256
- **Nested Loops with Control Statements** ([ch05_nested_loops_with_control_statements](./go-beyond-example-code/ch05_nested_loops_with_control_statements/)) - Line 1353
- **Performance Considerations for Nested Loops** ([ch05_performance_considerations_for_nested_loops](./go-beyond-example-code/ch05_performance_considerations_for_nested_loops/)) - Line 1452
- **Examples** ([ch05_examples_deb40b](./go-beyond-example-code/ch05_examples_deb40b/)) - Line 1510
- **Factorial Calculator** ([ch05_factorial_calculator](./go-beyond-example-code/ch05_factorial_calculator/)) - Line 1568
- **Prime Number Generator** ([ch05_prime_number_generator](./go-beyond-example-code/ch05_prime_number_generator/)) - Line 1617
- **🎮 Interactive Menu System** ([ch05_interactive_menu_system](./go-beyond-example-code/ch05_interactive_menu_system/)) - Line 1702
- **Exercise 1** ([ch05_exercise_1](./go-beyond-example-code/ch05_exercise_1/)) - Line 1910
- **Exercise 1** ([ch05_exercise_1_88b1de](./go-beyond-example-code/ch05_exercise_1_88b1de/)) - Line 1932
- **Exercise 3** ([ch05_exercise_3](./go-beyond-example-code/ch05_exercise_3/)) - Line 1961

### chapter6.md

- **Array Declaration and Initialization** ([ch06_array_declaration_and_initialization](./go-beyond-example-code/ch06_array_declaration_and_initialization/)) - Line 21
- **Array Types and Zero Values** ([ch06_array_types_and_zero_values](./go-beyond-example-code/ch06_array_types_and_zero_values/)) - Line 58
- **Example |** ([ch06_example](./go-beyond-example-code/ch06_example/)) - Line 97
- **Array Iteration Patterns** ([ch06_array_iteration_patterns](./go-beyond-example-code/ch06_array_iteration_patterns/)) - Line 133
- **Array Operations and Functions** ([ch06_array_operations_and_functions](./go-beyond-example-code/ch06_array_operations_and_functions/)) - Line 175
- **Slice Declaration and Initialization** ([ch06_slice_declaration_and_initialization](./go-beyond-example-code/ch06_slice_declaration_and_initialization/)) - Line 251
- **Understanding Slice Internals** ([ch06_understanding_slice_internals](./go-beyond-example-code/ch06_understanding_slice_internals/)) - Line 288
- **Slice Creation Patterns** ([ch06_slice_creation_patterns](./go-beyond-example-code/ch06_slice_creation_patterns/)) - Line 323
- **Example |** ([ch06_example_34d6bd](./go-beyond-example-code/ch06_example_34d6bd/)) - Line 372
- **Advanced Append Patterns** ([ch06_advanced_append_patterns](./go-beyond-example-code/ch06_advanced_append_patterns/)) - Line 403
- **Append Performance Considerations** ([ch06_append_performance_considerations](./go-beyond-example-code/ch06_append_performance_considerations/)) - Line 454
- **Slicing Operations** ([ch06_slicing_operations](./go-beyond-example-code/ch06_slicing_operations/)) - Line 496
- **Advanced Slicing Patterns** ([ch06_advanced_slicing_patterns](./go-beyond-example-code/ch06_advanced_slicing_patterns/)) - Line 528
- **Slice Manipulation Functions** ([ch06_slice_manipulation_functions](./go-beyond-example-code/ch06_slice_manipulation_functions/)) - Line 586
- **Copying Slices** ([ch06_copying_slices](./go-beyond-example-code/ch06_copying_slices/)) - Line 659
- **Advanced Copying Patterns** ([ch06_advanced_copying_patterns](./go-beyond-example-code/ch06_advanced_copying_patterns/)) - Line 691
- **Copy Performance and Best Practices** ([ch06_copy_performance_and_best_practices](./go-beyond-example-code/ch06_copy_performance_and_best_practices/)) - Line 746
- **Examples** ([ch06_examples](./go-beyond-example-code/ch06_examples/)) - Line 821
- **Shopping Cart** ([ch06_shopping_cart](./go-beyond-example-code/ch06_shopping_cart/)) - Line 995
- **Exercise 1** ([ch06_exercise_1](./go-beyond-example-code/ch06_exercise_1/)) - Line 1190
- **Exercise 2** ([ch06_exercise_2](./go-beyond-example-code/ch06_exercise_2/)) - Line 1226
- **Exercise 3** ([ch06_exercise_3](./go-beyond-example-code/ch06_exercise_3/)) - Line 1252
- **Memory Allocation Patterns** ([ch06_memory_allocation_patterns](./go-beyond-example-code/ch06_memory_allocation_patterns/)) - Line 1339
- **Performance Optimization Tips** ([ch06_performance_optimization_tips](./go-beyond-example-code/ch06_performance_optimization_tips/)) - Line 1417
- **Memory Management Best Practices** ([ch06_memory_management_best_practices](./go-beyond-example-code/ch06_memory_management_best_practices/)) - Line 1528

### chapter7.md

- **Basic Struct Definition** ([ch07_basic_struct_definition](./go-beyond-example-code/ch07_basic_struct_definition/)) - Line 32
- **example.com Address** ([ch07_examplecom_address](./go-beyond-example-code/ch07_examplecom_address/)) - Line 79
- **Struct Initialization** ([ch07_struct_initialization](./go-beyond-example-code/ch07_struct_initialization/)) - Line 115
- **Basic Methods** ([ch07_basic_methods](./go-beyond-example-code/ch07_basic_methods/)) - Line 221
- **Example** ([ch07_example](./go-beyond-example-code/ch07_example/)) - Line 315
- **Example** ([ch07_example_04b92f](./go-beyond-example-code/ch07_example_04b92f/)) - Line 468
- **Interface Satisfaction Rules** ([ch07_interface_satisfaction_rules](./go-beyond-example-code/ch07_interface_satisfaction_rules/)) - Line 581
- **Interface Composition** ([ch07_interface_composition](./go-beyond-example-code/ch07_interface_composition/)) - Line 660
- **Examples** ([ch07_examples](./go-beyond-example-code/ch07_examples/)) - Line 789
- **Basic Embedding** ([ch07_basic_embedding](./go-beyond-example-code/ch07_basic_embedding/)) - Line 935
- **Method Resolution Order** ([ch07_method_resolution_order](./go-beyond-example-code/ch07_method_resolution_order/)) - Line 1075
- **Interface Embedding** ([ch07_interface_embedding](./go-beyond-example-code/ch07_interface_embedding/)) - Line 1155
- **Multiple Embedding** ([ch07_multiple_embedding](./go-beyond-example-code/ch07_multiple_embedding/)) - Line 1279
- **Examples** ([ch07_examples_5b613f](./go-beyond-example-code/ch07_examples_5b613f/)) - Line 1380
- **example demonstrates a more sophisticated library system with different types of items, search capabilities, and user management.** ([ch07_example_demonstrates_a_more_sophisticated_library_](./go-beyond-example-code/ch07_example_demonstrates_a_more_sophisticated_library_/)) - Line 1597
- **Exercise 1** ([ch07_exercise_1](./go-beyond-example-code/ch07_exercise_1/)) - Line 1992
- **Exercise 2** ([ch07_exercise_2](./go-beyond-example-code/ch07_exercise_2/)) - Line 2150
- **Exercise 3** ([ch07_exercise_3](./go-beyond-example-code/ch07_exercise_3/)) - Line 2357
- **Basic Struct Embedding** ([ch07_basic_struct_embedding](./go-beyond-example-code/ch07_basic_struct_embedding/)) - Line 2551
- **Method Overriding** ([ch07_method_overriding](./go-beyond-example-code/ch07_method_overriding/)) - Line 2602
- **Interface Embedding** ([ch07_interface_embedding_e56b74](./go-beyond-example-code/ch07_interface_embedding_e56b74/)) - Line 2655

### chapter8.md

- **Basic Pointer Operations** ([ch08_basic_pointer_operations](./go-beyond-example-code/ch08_basic_pointer_operations/)) - Line 31
- **Memory Layout Visualization** ([ch08_memory_layout_visualization](./go-beyond-example-code/ch08_memory_layout_visualization/)) - Line 70
- **Nil Pointers** ([ch08_nil_pointers](./go-beyond-example-code/ch08_nil_pointers/)) - Line 103
- **Example |** ([ch08_example](./go-beyond-example-code/ch08_example/)) - Line 145
- **Pointer Arithmetic (Limited in Go)** ([ch08_pointer_arithmetic_limited_in_go](./go-beyond-example-code/ch08_pointer_arithmetic_limited_in_go/)) - Line 191
- **Multiple Pointers to Same Variable** ([ch08_multiple_pointers_to_same_variable](./go-beyond-example-code/ch08_multiple_pointers_to_same_variable/)) - Line 224
- **Passing by Value vs Reference** ([ch08_passing_by_value_vs_reference](./go-beyond-example-code/ch08_passing_by_value_vs_reference/)) - Line 261
- **Pointer Parameters** ([ch08_pointer_parameters](./go-beyond-example-code/ch08_pointer_parameters/)) - Line 304
- **Pointer Receivers** ([ch08_pointer_receivers](./go-beyond-example-code/ch08_pointer_receivers/)) - Line 335
- **Basic Error Handling** ([ch08_basic_error_handling](./go-beyond-example-code/ch08_basic_error_handling/)) - Line 407
- **Common Error Patterns** ([ch08_common_error_patterns](./go-beyond-example-code/ch08_common_error_patterns/)) - Line 446
- **example.com")** ([ch08_examplecom](./go-beyond-example-code/ch08_examplecom/)) - Line 519
- **Custom Error Types** ([ch08_custom_error_types](./go-beyond-example-code/ch08_custom_error_types/)) - Line 572
- **Advanced Custom Error Types** ([ch08_advanced_custom_error_types](./go-beyond-example-code/ch08_advanced_custom_error_types/)) - Line 632
- **Error Handling Patterns** ([ch08_error_handling_patterns](./go-beyond-example-code/ch08_error_handling_patterns/)) - Line 803
- **Examples** ([ch08_examples](./go-beyond-example-code/ch08_examples/)) - Line 933
- **Examples** ([ch08_examples_f64978](./go-beyond-example-code/ch08_examples_f64978/)) - Line 954
- **Example()** ([ch08_example_604eed](./go-beyond-example-code/ch08_example_604eed/)) - Line 1011
- **Advanced Panic and Recovery** ([ch08_advanced_panic_and_recovery](./go-beyond-example-code/ch08_advanced_panic_and_recovery/)) - Line 1102
- **Example()** ([ch08_example_d493ad](./go-beyond-example-code/ch08_example_d493ad/)) - Line 1240
- **Examples** ([ch08_examples_8bbfb0](./go-beyond-example-code/ch08_examples_8bbfb0/)) - Line 1351
- **Advanced Banking System with Pointers and Errors** ([ch08_advanced_banking_system_with_pointers_and_errors](./go-beyond-example-code/ch08_advanced_banking_system_with_pointers_and_errors/)) - Line 1429
- **File Operations with Error Handling** ([ch08_file_operations_with_error_handling](./go-beyond-example-code/ch08_file_operations_with_error_handling/)) - Line 1722
- **Advanced File System with Pointers and Error Handling** ([ch08_advanced_file_system_with_pointers_and_error_handl](./go-beyond-example-code/ch08_advanced_file_system_with_pointers_and_error_handl/)) - Line 1774
- **Exercise 1** ([ch08_exercise_1](./go-beyond-example-code/ch08_exercise_1/)) - Line 2150
- **Exercise 2** ([ch08_exercise_2](./go-beyond-example-code/ch08_exercise_2/)) - Line 2411
- **Exercise 3** ([ch08_exercise_3](./go-beyond-example-code/ch08_exercise_3/)) - Line 2685

### chapter9.md

- **Basic Map Declaration** ([ch09_basic_map_declaration](./go-beyond-example-code/ch09_basic_map_declaration/)) - Line 45
- **Map Initialization** ([ch09_map_initialization](./go-beyond-example-code/ch09_map_initialization/)) - Line 74
- **Basic Access** ([ch09_basic_access](./go-beyond-example-code/ch09_basic_access/)) - Line 119
- **Checking Key Existence** ([ch09_checking_key_existence](./go-beyond-example-code/ch09_checking_key_existence/)) - Line 164
- **Adding and Updating** ([ch09_adding_and_updating](./go-beyond-example-code/ch09_adding_and_updating/)) - Line 230
- **Deleting Elements** ([ch09_deleting_elements](./go-beyond-example-code/ch09_deleting_elements/)) - Line 263
- **Basic Iteration** ([ch09_basic_iteration](./go-beyond-example-code/ch09_basic_iteration/)) - Line 300
- **Ordered Iteration** ([ch09_ordered_iteration](./go-beyond-example-code/ch09_ordered_iteration/)) - Line 335
- **Examples** ([ch09_examples](./go-beyond-example-code/ch09_examples/)) - Line 372
- **Word Frequency Counter** ([ch09_word_frequency_counter](./go-beyond-example-code/ch09_word_frequency_counter/)) - Line 420
- **main_program** ([ch09_main_program](./go-beyond-example-code/ch09_main_program/)) - Line 481
- **Contact Book** ([ch09_contact_book](./go-beyond-example-code/ch09_contact_book/)) - Line 522
- **main_program** ([ch09_main_program_8f3b56](./go-beyond-example-code/ch09_main_program_8f3b56/)) - Line 614
- **example.com",** ([ch09_examplecom](./go-beyond-example-code/ch09_examplecom/)) - Line 698
- **Map Merging** ([ch09_map_merging](./go-beyond-example-code/ch09_map_merging/)) - Line 759
- **Nested Maps** ([ch09_nested_maps](./go-beyond-example-code/ch09_nested_maps/)) - Line 846
- **Exercise 1** ([ch09_exercise_1](./go-beyond-example-code/ch09_exercise_1/)) - Line 938
- **Exercise 1** ([ch09_exercise_1_b2a212](./go-beyond-example-code/ch09_exercise_1_b2a212/)) - Line 961
- **Exercise 3** ([ch09_exercise_3](./go-beyond-example-code/ch09_exercise_3/)) - Line 995


## Usage

Each program directory contains:
- `main.go` - The main program file
- `go.mod` - Go module definition
- `README.md` - Program description and usage

To run any program:

```bash
cd <program-directory>
go run main.go
```

To build any program:

```bash
cd <program-directory>
go build -o program main.go
./go-beyond-example-code/program
```

## Testing All Programs

To test that all programs compile successfully:

```bash
find . -name "main.go" -execdir go build -o /tmp/test_build {} \;
```

Generated on: Sat Sep 27 14:07:33 PDT 2025
Total programs: 444
