# Contributing to hxgo

## Our Development Philosophy

hxgo is built with a specific architectural vision and performance constraints. We use `unsafe.Pointer` deliberately and maintain strict memory safety guarantees through careful design patterns. This approach requires good understanding of Go's memory model and escape analysis.

## How to Contribute Effectively

### For Feature Requests and Improvements

We welcome ideas and improvements, but please **start with discussion before code**:

1. **Open an Issue First**: Describe the problem or enhancement
2. **Articulate the Use Case**: Help us understand the real-world need
3. **Discuss Design Impact**: How does this fit with hxgo's constraints?
4. **Collaborate**: We will work with you on implementation approach

### Why Discussion First?

- hxgo uses `unsafe` patterns that require careful review
- Memory safety depends on specific usage patterns
- We maintain strict performance guarantees
- Architectural consistency is crucial

## Contribution Areas Where We Need Help

### Lower-Risk Contributions (Great Starting Points)

- **Documentation improvements**
- **Example applications and tutorials**
- **Bug reports with reproduction cases**
- **Performance benchmark additions**
- **Test coverage expansion**

### Higher-Risk Contributions (Require Deep Review)

- **Core renderer changes**
- **Memory management modifications**
- `unsafe.Pointer` usage changes
- **Architectural modifications**

## Our Review Process

Given hxgo's performance-sensitive nature, we have a thorough review process:

1. **Design Review**: We'll discuss the approach before any code
2. **Safety Audit**: All `unsafe` usage gets extra scrutiny
3. **Performance Validation**: Benchmarks must show no regression
4. **Integration Testing**: Must work within existing constraints

## Alternative Contribution Paths

If you are uncomfortable with hxgo's `unsafe` patterns but want to help:

1. **Build with hxgo**: Create applications that showcase its capabilities
2. **Write tutorials**: Help others learn the patterns
3. **Report issues**: Detailed bug reports are invaluable
4. **Suggest improvements**: Your user perspective helps shape the roadmap

## AGPL and Sharing Improvements

hxgo uses the AGPL-3.0 license, which means:

- You can use, modify, and distribute hxgo
- If you modify hxgo and provide it as a service, you must share your changes
- This ensures improvements benefit the entire community

We believe this creates a fair ecosystem where everyone benefits from collective innovation.

## Safety First

Remember: hxgo's memory safety depends on specific usage patterns. If you are proposing changes that touch:

- `accumulator` internals
- `unsafe.Pointer` usage  
- Node lifetime management
- Renderer core logic

Please be prepared for detailed discussion and multiple review iterations. Our priority is maintaining safety while delivering exceptional performance.  
