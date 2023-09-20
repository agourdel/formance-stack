# PostTransactionScript


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      | Example                                                                          |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `plain`                                                                          | *str*                                                                            | :heavy_check_mark:                                                               | N/A                                                                              | vars {<br/>account $user<br/>}<br/>send [COIN 10] (<br/>	source = @world<br/>	destination = $user<br/>)<br/> |
| `vars`                                                                           | dict[str, *Any*]                                                                 | :heavy_minus_sign:                                                               | N/A                                                                              |                                                                                  |