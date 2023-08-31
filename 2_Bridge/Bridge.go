package main

import (
	"fmt"
	"io"
	"math/rand"
)

// IDataExporter 是数据导出器接口
type IDataExporter interface {
	Fetcher(fetcher IDataFetcher)              // 设置数据查询器
	Export(sql string, writer io.Writer) error // 导出数据
}

// IDataFetcher 是数据查询器接口
type IDataFetcher interface {
	Fetch(sql string) []interface{} // 查询数据
}

// NewMysqlDataFetcher 创建一个新的 MysqlDataFetcher 实例
func NewMysqlDataFetcher(configStr string) IDataFetcher {
	return &MysqlDataFetcher{
		Config: configStr,
	}
}

// MysqlDataFetcher 是 Mysql 数据查询器
type MysqlDataFetcher struct {
	Config string // 配置信息
}

// Fetch 从 Mysql 数据源查询数据
func (mf *MysqlDataFetcher) Fetch(sql string) []interface{} {
	fmt.Println("Fetch data from mysql source: " + mf.Config)
	rows := make([]interface{}, 0)
	// 插入两个随机数组成的切片，模拟查询要返回的数据集
	rows = append(rows, rand.Perm(10), rand.Perm(10))
	return rows
}

// NewOracleDataFetcher 创建一个新的 OracleDataFetcher 实例
func NewOracleDataFetcher(configStr string) IDataFetcher {
	return &OracleDataFetcher{
		Config: configStr,
	}
}

// OracleDataFetcher 是 Oracle 数据查询器
type OracleDataFetcher struct {
	Config string // 配置信息
}

// Fetch 从 Oracle 数据源查询数据
func (of *OracleDataFetcher) Fetch(sql string) []interface{} {
	fmt.Println("Fetch data from oracle source: " + of.Config)
	rows := make([]interface{}, 0)
	// 插入两个随机数组成的切片，模拟查询要返回的数据集
	rows = append(rows, rand.Perm(10), rand.Perm(10))
	return rows
}

// NewCsvExporter 创建一个新的 CsvExporter 实例
func NewCsvExporter(fetcher IDataFetcher) IDataExporter {
	return &CsvExporter{
		mFetcher: fetcher,
	}
}

// CsvExporter 是 CSV 数据导出器
type CsvExporter struct {
	mFetcher IDataFetcher // 数据查询器
}

// Fetcher 设置数据查询器
func (ce *CsvExporter) Fetcher(fetcher IDataFetcher) {
	ce.mFetcher = fetcher
}

// Export 导出数据到 CSV
func (ce *CsvExporter) Export(sql string, writer io.Writer) error {
	rows := ce.mFetcher.Fetch(sql)
	fmt.Printf("CsvExporter.Export, got %v rows\n", len(rows))
	for i, v := range rows {
		fmt.Printf("  行号: %d 值: %s\n", i+1, v)
	}
	return nil
}

// NewJsonExporter 创建一个新的 JsonExporter 实例
func NewJsonExporter(fetcher IDataFetcher) IDataExporter {
	return &JsonExporter{
		mFetcher: fetcher,
	}
}

// JsonExporter 是 JSON 数据导出器
type JsonExporter struct {
	mFetcher IDataFetcher // 数据查询器
}

// Fetcher 设置数据查询器
func (je *JsonExporter) Fetcher(fetcher IDataFetcher) {
	je.mFetcher = fetcher
}

// Export 导出数据到 JSON
func (je *JsonExporter) Export(sql string, writer io.Writer) error {
	rows := je.mFetcher.Fetch(sql)
	fmt.Printf("JsonExporter.Export, got %v rows\n", len(rows))
	for i, v := range rows {
		fmt.Printf("  行号: %d 值: %s\n", i+1, v)
	}
	return nil
}
