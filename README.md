* 数据结构

type SelpgArgs struct {
    
    	startP int    //start page number
	endP   int    //end page number
	len    int    //how many lines per page
	typeP  int    //type of page decide whether to yse '/f'
	desP   string //destination to print
	name   string //file name
}

* 函数声明

func Processor1(sa SelpgArgs, f bool, p bool)

func Processor2(sa SelpgArgs, f bool, p bool)

func usage() 

* 结果测试

生成两个测试文档：a.txt b.txt. a.txt以\n结尾，b.txt以\f结尾。两个文档均为1-100的整数数字


1.见Screenshots1.1 1.2
 
      显示1-72个整数
  
2.见Screenshots2 

      结果和1一样
  
4.见Screenshots4

     a_out.txt被写入1-72的整数
  
5.见Screenshots5.1 5.2

     屏幕无输出
    
     errorF文件被写入了错误信息


6.见Screenshots6 6.1 6.2

    errorF文件被写入了错误信息

    a_out写入了错误前的输出

    屏幕无输出
  

7.见Screenshots7.1 7.2

     和6一样：屏幕无输出，a_out有输出
  
8.见Screenshots8

     屏幕显示错误信息
  
9.见Screenshots9.1 9.2

      go代码已给出
      屏幕输出1-72整数
      
10.见Screenshots10.1 10.2

      命令行无错误提示
      errorF显示错误信息
