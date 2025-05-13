# go-core

{
    "title": "Container Read Quota",
    "type": "query_table",
    "requests": [
        {
            "queries": [
                {
                    "data_source": "metrics",
                    "name": "query6",
                    "query": "max:kvsapi.container.rate_limit_throughput.read{container:*recomm_items_api} by {container,segment,cluster}.as_count().fill(last).rollup(max, 300)",
                    "aggregator": "max"
                },
                {
                    "query": "sum:kvsapi.http.effective.request{method:get,!engine:search,container:*recomm_items_api} by {container,segment,cluster}.as_count().rollup(sum, 60)",
                    "data_source": "metrics",
                    "name": "query1",
                    "aggregator": "max"
                },
                {
                    "query": "sum:kvsapi.http.effective.request{method:post,!engine:search,action:bulk_get,container:*recomm_items_api} by {container,segment,cluster}.as_count().rollup(sum, 60)",
                    "data_source": "metrics",
                    "name": "query2",
                    "aggregator": "max"
                },
                {
                    "query": "sum:kvsapi.http.effective.request{method:post,!engine:search,action:batch_get,container:*recomm_items_api} by {container,segment,cluster}.as_count().rollup(sum, 60)",
                    "data_source": "metrics",
                    "name": "query3",
                    "aggregator": "max"
                },
                {
                    "query": "sum:kvsapi.http.effective.request{method:post,!engine:search,action:bulk_read,container:*recomm_items_api} by {container,segment,cluster}.as_count().rollup(sum, 60)",
                    "data_source": "metrics",
                    "name": "query4",
                    "aggregator": "max"
                },
                {
                    "query": "sum:kvsapi.http.effective.request{method:post,!engine:search,action:batch_read,container:*recomm_items_api} by {container,segment,cluster}.as_count().rollup(sum, 60)",
                    "data_source": "metrics",
                    "name": "query5",
                    "aggregator": "max"
                },
                {
                    "data_source": "metrics",
                    "name": "query11",
                    "query": "sum:aws.dynamodb.provisioned_read_capacity_units{container:*recomm_items_api} by {container,segment,cluster}.rollup(sum, 60)",
                    "aggregator": "max"
                },
                {
                    "data_source": "metrics",
                    "name": "query7",
                    "query": "sum:kvsapi.http.effective.request{method:get,!engine:search,container:*recomm_items_api} by {container,segment,cluster}.as_count().rollup(sum, 60)",
                    "aggregator": "max"
                },
                {
                    "data_source": "metrics",
                    "name": "query8",
                    "query": "sum:kvsapi.http.effective.request{method:post,!engine:search,action:bulk_get,container:*recomm_items_api} by {container,segment,cluster}.as_count().rollup(sum, 60)",
                    "aggregator": "max"
                },
                {
                    "data_source": "metrics",
                    "name": "query9",
                    "query": "sum:kvsapi.http.effective.request{method:post,!engine:search,action:batch_get,container:*recomm_items_api} by {container,segment,cluster}.as_count().rollup(sum, 60)",
                    "aggregator": "max"
                },
                {
                    "data_source": "metrics",
                    "name": "query12",
                    "query": "sum:kvsapi.http.effective.request{method:post,!engine:search,action:bulk_read,container:*recomm_items_api} by {container,segment,cluster}.as_count().rollup(sum, 60)",
                    "aggregator": "max"
                },
                {
                    "data_source": "metrics",
                    "name": "query10",
                    "query": "sum:aws.dynamodb.consumed_read_capacity_units{container:*recomm_items_api} by {container,segment,cluster}.rollup(sum, 60)",
                    "aggregator": "max"
                },
                {
                    "data_source": "metrics",
                    "name": "query13",
                    "query": "sum:kvsapi.http.effective.request{method:post,!engine:search,action:batch_read,container:*recomm_items_api} by {container,segment,cluster}.as_count().rollup(sum, 60)",
                    "aggregator": "max"
                }
            ],
            "response_format": "scalar",
            "sort": {
                "count": 500,
                "order_by": [
                    {
                        "type": "formula",
                        "index": 3,
                        "order": "desc"
                    }
                ]
            },
            "formulas": [
                {
                    "alias": "Max",
                    "cell_display_mode": "number",
                    "formula": "query6"
                },
                {
                    "alias": "Consumed",
                    "cell_display_mode": "trend",
                    "cell_display_mode_options": {
                        "trend_type": "area",
                        "y_scale": "independent"
                    },
                    "formula": "query1 + query2 + query3 + query4 + query5"
                },
                {
                    "alias": "Provisioned",
                    "cell_display_mode": "trend",
                    "cell_display_mode_options": {
                        "trend_type": "area",
                        "y_scale": "independent"
                    },
                    "formula": "query11 * (query7 + query8 + query9 + query12 + query12) / query10"
                },
                {
                    "alias": "% Consumed from Max",
                    "cell_display_mode": "trend",
                    "cell_display_mode_options": {
                        "trend_type": "line",
                        "y_scale": "shared"
                    },
                    "formula": "(query1 + query2 + query3 + query4 + query5) / query6 * 100"
                }
            ]
        }
    ],
    "has_search_bar": "auto"
}
