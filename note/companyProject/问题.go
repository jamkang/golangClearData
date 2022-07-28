package companyProject

/*
1、table中handlerDealBet(player *Player, msg *proto.StEmptyMsg){ 中
2、c++中的end问题，在server.go 的 GetRobotBeHavior 中
4、table中func (self *Table) SendNormalFace(seatid, type_ int) {最后一条语句
*/

/*
1、在table中
	func handlerTableLogin(
2、在msg——handler.go中
	func (self *Server) writeLoginLog(player *Player) { 发布订阅有些问题
	func (self *Server) handleChangeTable(player *Player, msg *proto.StEmptyMsg) {client
3、在player中
	StartHeartBeat和StopHeartBeat不会改写
4、在server.go中
	关于server和game中元素冲突事情，
	func (s *Server) subscribe(subdata string) {中packag的问题

5. 现在存在的bug
	1.前期使用redis的时候,没有判断是否为机器人
	2. glog以前我是使用logger这个来记录日志,
*/
