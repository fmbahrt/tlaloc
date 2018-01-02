package routing

import (
    "../handlers" // change this import
    "../stats"
)

// Perhaps move this
var stat stats.Stats

var ServerRoutes = Routes{
    Route{
        "Index",
        "GET",
        "/",
        handlers.Index,
    },
    Route{
        "Search",
        "GET",
        "/_search",
        stat.Decorate(handlers.Search),
    },
    Route{
        "Statistics",
        "GET",
        "/statz",
        stat.Statz,
    },
    Route{
        "ScatterPNG",
        "GET",
        "/statz/scatter.png",
        stat.Scatter,
    },
    Route{
        "HistPNG",
        "GET",
        "/statz/hist.png",
        stat.Hist,
    },
    Route{
        "Config",
        "GET",
        "/config",
        handlers.Config,
    },
}
