//simple for loop iterate from 1 to n
fn sum_to_n_a(n: i32) -> i32 {
    let mut sum = 0;
    for i in 1..=n {
        sum += i;
    }
    sum
}

//math formunla
fn sum_to_n_b(n: i32) -> i32 {
    n * (n + 1) / 2
}

//recursion
fn sum_to_n_c(n: i32) -> i32 {
    if n <= 1 {
        n
    } else {
        n + sum_to_n_c(n - 1)
    }
}

fn main() {
    println!("{}",sum_to_n_a(5));
    println!("{}",sum_to_n_b(5));
    println!("{}",sum_to_n_c(5));
}
