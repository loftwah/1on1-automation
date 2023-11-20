resource "aws_security_group" "1on1_playground_sg" {
  name        = "1on1-automation-sg"
  description = "Security group for 1on1 Playground"
  vpc_id      = data.aws_vpc.default.id

  ingress {
    from_port   = 1323
    to_port     = 1323
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_ecs_cluster" "1on1_playground" {
  name = "1on1-automation"
}

resource "aws_ecs_task_definition" "1on1_task" {
  family                   = "1on1-automation"
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  execution_role_arn       = "arn:aws:iam::614442032670:role/ecsTaskExecutionRole"
  cpu                      = "256"
  memory                   = "512"

  container_definitions = jsonencode([{
    name  = "1on1-automation",
    image = "614442032670.dkr.ecr.ap-southeast-2.amazonaws.com/1on1-automation:latest",
    cpu   = 128,
    memory = 256,
    essential = true,
    portMappings = [{
      containerPort = 1323
    }]
  }])
}

resource "aws_ecs_service" "1on1_service" {
  name            = "1on1-automation"
  cluster         = aws_ecs_cluster.1on1_playground.id
  task_definition = aws_ecs_task_definition.1on1_task.arn
  launch_type     = "FARGATE"
  desired_count   = 1

  network_configuration {
    subnets         = data.aws_subnets.default.ids
    security_groups = [aws_security_group.1on1_playground_sg.id]
    assign_public_ip = true
  }

  load_balancer {
    target_group_arn = aws_lb_target_group.1on1_tg.arn
    container_name   = "1on1-automation"
    container_port   = 1323
  }

  depends_on = [
    aws_lb_listener.1on1_listener,
  ]
}

