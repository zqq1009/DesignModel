package main

import (
	"bytes"
	"testing"
)

func TestBridge(t *testing.T) {
	mFetcher := NewMysqlDataFetcher("mysql://127.0.0.1:3306")
	csvExporter := NewCsvExporter(mFetcher)
	var writer bytes.Buffer
	// 从 MySQL 数据源导出 CSV
	csvExporter.Export("select * from xxx", &writer)

	oFetcher := NewOracleDataFetcher("mysql://127.0.0.1:1001")
	csvExporter.Fetcher(oFetcher)
	// 从 Oracle 数据源导出 CSV
	csvExporter.Export("select * from xxx", &writer)

	// 从 MySQL 数据源导出 JSON
	jsonExporter := NewJsonExporter(mFetcher)
	jsonExporter.Export("select * from xxx", &writer)
}
