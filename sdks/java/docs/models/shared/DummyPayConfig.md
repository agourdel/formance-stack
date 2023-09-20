# DummyPayConfig


## Fields

| Field                                                                                         | Type                                                                                          | Required                                                                                      | Description                                                                                   | Example                                                                                       |
| --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- |
| `directory`                                                                                   | *String*                                                                                      | :heavy_check_mark:                                                                            | N/A                                                                                           | /tmp/dummypay                                                                                 |
| `fileGenerationPeriod`                                                                        | *String*                                                                                      | :heavy_minus_sign:                                                                            | The frequency at which the connector will create new payment objects in the directory         | 60s                                                                                           |
| `filePollingPeriod`                                                                           | *String*                                                                                      | :heavy_minus_sign:                                                                            | The frequency at which the connector will try to fetch new payment objects from the directory | 60s                                                                                           |