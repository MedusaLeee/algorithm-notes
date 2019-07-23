package cardSort

/**
for (int i = 0; i < 52; ++i) {
		//insert to proper place
		int j;
		for (j = 0; j < i; ++j) {//在本轮范围之内
			if (c.first() > c.second()) {
				c.swap();//将大的向下移动至合适的位置
				c.first_go_down();//将小的pass
				continue;
			}
			break;
		}
		for (int k = j; k < 51; ++k) {
			c.first_go_down();//本轮排序完毕，将位置排列好，置于下轮开始状态
		}
	}
	//最后差一次
	c.first_go_down();
---------------------
*/
type CardSort struct {
	arr []int
}

func (cs *CardSort) Sort() []int {
	length := len(cs.arr)
	for i := 1; i <= length-1; i++ {
		//第i轮操作需要n-i次换牌操作，以及i次将最上面的牌移到最下面
		for j := 0; i < length-i; j++ {
			if cs.arr[1] < cs.arr[0] {
				cs.swap(0, 1)
			}
			cs.toLast(1)
		}
		cs.toLast(i)
	}
	return cs.arr
}

func (cs *CardSort) swap(j int, i int) {
	temp := cs.arr[j]
	cs.arr[j] = cs.arr[i]
	cs.arr[i] = temp
}

func (cs *CardSort) toLast(num int) {
	length := len(cs.arr)
	for i := 0; i < num; i++ {
		for j := 0; j < length-1; j++ {
			cs.swap(j, j+1)
		}
	}
}
