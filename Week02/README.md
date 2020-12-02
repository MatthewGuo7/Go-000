我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

答案：
不应该直接抛sql.ErrNoRows. 原因：因为这是dao层，若需求变动，迁移到其他数据库， 或增加缓存，则肯定不是判断sql.ErrNoRows.
解决方案：
在repos层或者dao层封装一个自己应用层的sentinal error.遇到 sql.ErrNoRows的时候，转换一下，抛这个sentinal error
